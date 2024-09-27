(() => {

    let fatherDir;
    let pageFiles;

    main();

    function main() {
        document.getElementById('sort-btn').addEventListener('click', showSortDialog);
        document.querySelector('#sort-dialog .select-widget select.by').addEventListener('change', sortBySelectEvent);
        document.querySelector('#sort-dialog .btn-group .submit').addEventListener('click', submitSortSelectEvent);
        document.querySelector('#sort-dialog .btn-group .close').addEventListener('click', closeSortDialog);

        // ajax.getJson(`/api/index${location.search}`, (response) => {
        //     fatherDir = response['father'];
        //     pageFiles = response['files'];
        //     pageFiles = sortFiles(pageFiles);
        //
        //     renderContent(fatherDir, pageFiles);
        // });

        // region 测试用代码
        const response = {
            "father": "/root/file-server-android",
            "files": [
                {"filename": "ab.mp3", "is_dir": false, "size": 24576},
                {"filename": "frontend", "is_dir": true, "size": 0},
                {"filename": "database", "is_dir": true, "size": 2049},
                {"filename": ".env", "is_dir": false, "size": 59},
                {"filename": "filename", "is_dir": false, "size": 59},
                {"filename": "filename.txt", "is_dir": false, "size": 100},
                {"filename": "filename.zip", "is_dir": false, "size": 2457622},
                {"filename": "filename.mp4", "is_dir": false, "size": 245761024},
                {"filename": "655.jpg", "is_dir": false, "size": 24576}
            ]
        }

        fatherDir = response['father'];
        pageFiles = response['files'];
        pageFiles = sortFiles(pageFiles);

        renderContent(fatherDir, pageFiles);
        // endregion
    }

    function renderContent(fatherDir, files) {
        if (files.length === 0) {
            return undefined;
        }

        const searchParams = parseSearchParams();
        const visibleDir = searchParams['visible_dir'];
        const path = searchParams['path'];
        let hrefTemplate;
        if (visibleDir === undefined && path === undefined) {
            // 场景1
            document.getElementById('sort-btn').classList.add('hidden');
            document.querySelector('#file_list .up-level').classList.add('hidden');
            document.getElementById('current-location').classList.add('hidden');
            hrefTemplate = (filename) => { return `/?visible_dir=${filename}` };
        } else if (visibleDir !== undefined && path === undefined) {
            // 场景2
            document.querySelector('#file_list .up-level').href = '/';
            document.getElementById('current-location').textContent = fatherDir;
            hrefTemplate = (filename) => { return `/?visible_dir=${fatherDir}&path=${filename}` };
        } else if (visibleDir !== undefined && path !== undefined) {
            // 场景3
            const isMultiLevelPath = inArray('/', path.replace(/\\/g, '/'));
            const href1 = `/?visible_dir=${visibleDir}&path=${getUpLevelPath(path)}`;
            const href2 = `/?visible_dir=${visibleDir}`;
            document.querySelector('#file_list .up-level').href = isMultiLevelPath ? href1 : href2;
            document.getElementById('current-location').textContent = fatherDir;
            hrefTemplate = (filename) => { return `/?visible_dir=${visibleDir}&path=${connectPath(path, filename)}` };
        } else {
            return undefined;
        }

        const fileList = document.getElementById('file_list');

        fileList.querySelector('.empty-hint-text').classList.add('hidden');
        fileList.querySelectorAll('.file').forEach((f) => {
            f.remove();
        });

        for (let i = 0; i < files.length; i++) {
            let info = files[i];
            let href = hrefTemplate(info['filename']);
            let fileType = info['is_dir'] ? 'dir' : getFileTypeFromExtension(info['filename']);
            let fileSize = info['is_dir'] ? '' : getHumanFileSize(info['size']);
            let html = `
                <a class="file" href="${href}">
                    ${fileTypeIconMap[fileType]}
                    <div class="filename">${info['filename']}</div>
                    <div class="filesize">${fileSize}</div>
                </a>`;
            let ele = htmlToElement(html);
            ele.querySelector('svg').setAttribute('class', `icon icon-${fileType}`);

            fileList.appendChild(ele);
        }
    }

    function sortFilesByType(files, reversed) {
        const filesObject = {};
        for (let i = 0; i < fileTypeOrder.length; i++) {
            filesObject[fileTypeOrder[i]] = [];
        }
        for (let i = 0; i < files.length; i++) {
            let info = files[i];
            let k = info['is_dir'] ? 'dir' : getFileTypeFromExtension(info['filename']);
            filesObject[k].push(info);
        }
        for (let i = 0; i < fileTypeOrder.length; i++) {
            let k = fileTypeOrder[i];
            filesObject[k] = sortFilesByName(filesObject[k], false);
        }

        const mergeKeys = [];
        for (let i = 0; i < fileTypeOrder.length; i++) {
            mergeKeys.push(fileTypeOrder[i]);
        }
        if (reversed === true) {
            mergeKeys.reverse();
        }

        let result = [];
        for (let i = 0; i < mergeKeys.length; i++) {
            result = result.concat(filesObject[mergeKeys[i]]);
        }

        return result;
    }

    function sortFilesByName(files, reversed) {
        files.sort((a, b) => {
            const fn1 = a['filename'];
            const fn2 = b['filename'];
            const length = fn1.length > fn2.length ? fn2.length : fn1.length;
            for (let i = 0; i < length; i++) {
                let c1 = fn1.charCodeAt(i);
                let c2 = fn2.charCodeAt(i);
                if (c1 === c2) {
                    continue;
                }
                return c1 - c2;
            }

            if (fn1.length > fn2.length) {
                return 1;
            } else if (fn1.length === fn2.length) {
                return 0;
            } else {
                return -1;
            }
        });
        if (reversed === true) {
            files.reverse();
        }
        return files;
    }

    function sortFilesBySize(files, reversed) {
        let isDirFiles = [];
        const notDirFiles = [];
        for (let i = 0; i < files.length; i++) {
            if (files[i]['is_dir']) {
                isDirFiles.push(files[i]);
            } else {
                notDirFiles.push(files[i]);
            }
        }

        isDirFiles = sortFilesByName(isDirFiles, false);

        notDirFiles.sort((a, b) => {
            return a['size'] - b['size'];
        });
        if (reversed === true) {
            notDirFiles.reverse();
        }

        return isDirFiles.concat(notDirFiles);
    }

    function sortFiles(files, by, reversed) {
        if (by !== undefined && by !== 'type' && by !== 'name' && by !== 'size') {
            throw 'by参数不可用'
        }
        if (by === undefined) {
            by = localStorage.getItem('sort_files_by');
            by = by === null ? 'type' : by;
        } else {
            localStorage.setItem('sort_files_by', by);
        }

        if (reversed === undefined) {
            reversed = localStorage.getItem('sort_files_reversed') === 'true';
        } else {
            localStorage.setItem('sort_files_reversed', String(reversed));
        }

        if (by === 'type') {
            return sortFilesByType(files, reversed);
        } else if (by === 'name') {
            return sortFilesByName(files, reversed);
        } else if (by === 'size') {
            return sortFilesBySize(files, reversed);
        }
        return files;
    }

    function showSortDialog() {
        document.getElementById('overlay').classList.remove('hidden');
        document.getElementById('sort-dialog').classList.remove('hidden');
    }

    function closeSortDialog() {
        document.getElementById('overlay').classList.add('hidden');
        document.getElementById('sort-dialog').classList.add('hidden');
    }

    function sortBySelectEvent(e) {
        const select = e.target;
        const option = select.options[select.selectedIndex];
        const control = option.getAttribute('data-control');
        document.querySelector('#sort-dialog .select-widget select.reversed:not(.hidden)').classList.add('hidden');
        document.querySelector(`#sort-dialog .select-widget select.reversed[data-control="${control}"]`).classList.remove('hidden');
    }

    function submitSortSelectEvent(e) {
        const by = document.querySelector('#sort-dialog .select-widget select.by').value;
        const reversed = document.querySelector('#sort-dialog .select-widget select.reversed:not(.hidden)').value === '1';

        renderContent(fatherDir, sortFiles(pageFiles, by, reversed))

        closeSortDialog();
    }
})();
