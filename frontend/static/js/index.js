(() => {

    main();

    function main() {
        ajax.getJson(`/api/index${location.search}`, (response) => {
            renderContent(response['files']);
        });

        renderContent([
            'C:\\',
            'D:\\',
            'E:\\',
            'C:\\Users\\504-Jiyue-Xie\\Desktop\\multimodal_guiyi',
            'D:\\EvaluationExperiment',
            'D:\\桂医课程项目\\CPGGBDT',
            'D:\\文档\\PPT模板',
            'E:\\Music',
        ]);
    }

    function renderContent(filenames) {
        if (filenames.length === 0) {
            return undefined;
        }

        const fileList = document.getElementById('file_list');
        fileList.querySelector('.empty-hint-text').remove();

        for (let i = 0; i < filenames.length; i++) {
            let filename = filenames[i];
            let href = `/?visible_dir=${filenames[i]}`;
            let html = `
                <a class="file" href="${href}">
                    <svg class="icon" xmlns="http://www.w3.org/2000/svg" width="146.43" height="128" viewBox="0 0 1025 896"><path d="M960.232 896h-896q-26 0-45-19t-19-45V256h1024v576q0 26-19 45t-45 19m-960-704V64q0-27 18.5-45.5T64.232 0h384q26 0 45 18.5t19 45.5t18.5 45.5t45.5 18.5h384q26 0 45 19t19 45z"/></svg>
                    <div class="filename">${filename}</div>
                </a>`;

            fileList.appendChild(htmlToElement(html));
        }
    }
})();