window.extensionTypeMap = {
    'rar': 'package',
    'zip': 'package',
    '7z': 'package',
    'gz': 'package',
    'tar': 'package',
    'mp4': 'video',
    'm4v': 'video',
    'mkv': 'video',
    'webm': 'video',
    'mov': 'video',
    'avi': 'video',
    'wmv': 'video',
    'mpg': 'video',
    'flv': 'video',
    'mpeg': 'video',
    'rm': 'video',
    'ram': 'video',
    'rmvb': 'video',
    'jpg': 'image',
    'png': 'image',
    'jpeg': 'image',
    'gif': 'image',
    'webp': 'image',
    'ico': 'image',
    'bmp': 'image',
    'psd': 'image',
    'dwg': 'image',
    'xcf': 'image',
    'jpx': 'image',
    'apng': 'image',
    'cr2': 'image',
    'tif': 'image',
    'jxr': 'image',
    'heic': 'image',
    'mp3': 'audio',
    'wav': 'audio',
    'm4a': 'audio',
    'flac': 'audio',
    'aac': 'audio',
    'ogg': 'audio',
    'mid': 'audio',
    'amr': 'audio',
    'aiff': 'audio',
    'txt': 'text',
    'py': 'text',
    'js': 'text',
    'ipynb': 'text',
    'ini': 'text',
    'css': 'text',
    'scss': 'text',
    'sass': 'text',
    'html': 'text',
    'xml': 'text',
    'json': 'text',
    'java': 'text',
    'c': 'text',
    'cpp': 'text',
    'md': 'text'
}

window.fileTypeIconMap = {
    'dir': `<svg xmlns="http://www.w3.org/2000/svg" width="146.43" height="128" viewBox="0 0 1025 896"><path d="M960.232 896h-896q-26 0-45-19t-19-45V256h1024v576q0 26-19 45t-45 19m-960-704V64q0-27 18.5-45.5T64.232 0h384q26 0 45 18.5t19 45.5t18.5 45.5t45.5 18.5h384q26 0 45 19t19 45z"/></svg>`,
    'package': `<svg xmlns="http://www.w3.org/2000/svg" width="128.13" height="128" viewBox="0 0 1025 1024"><path d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h320v64h-32q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h32v64h-32q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h32v64h-32q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h32v64h-32q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h32v64q-27 0-45.5 19t-18.5 45v192q0 27 18.5 45.5t45.5 18.5h128q27 0 45.5-18.5t18.5-45.5V640q0-26-18.5-45t-45.5-19h32q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-32v-64h32q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-32v-64h32q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-32v-64h32q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-32V64h32q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h288q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5m-384-192q-26 0-45-18.5t-19-45t18.5-45.5t45.5-19t45.5 19t18.5 45.5t-19 45t-45 18.5"/></svg>`,
    'video': `<svg xmlns="http://www.w3.org/2000/svg" width="128" height="128" viewBox="0 0 48 48"><defs><mask id="IconifyId1922cdc16e4b745fe8"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 10a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v28a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-linecap="round" d="M36 8v32M12 8v32m26-22h6m-6 12h6M4 18h6m-6-2v4M9 8h6M9 40h6M33 8h6m-6 32h6M4 30h6m-6-2v4m40-4v4m0-16v4"/><path fill="#555" d="m21 19l8 5l-8 5z"/></g></mask></defs><path d="M0 0h48v48H0z" mask="url(#IconifyId1922cdc16e4b745fe8)"/></svg>`,
    'audio': `<svg xmlns="http://www.w3.org/2000/svg" width="128" height="128" viewBox="0 0 24 24"><path d="M16 9h-3v5.5a2.5 2.5 0 0 1-2.5 2.5A2.5 2.5 0 0 1 8 14.5a2.5 2.5 0 0 1 2.5-2.5c.57 0 1.08.19 1.5.5V7h4m3-4H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2"/></svg>`,
    'image': `<svg xmlns="http://www.w3.org/2000/svg" width="128" height="128" viewBox="0 0 32 32"><path d="M19 14a3 3 0 1 0-3-3a3 3 0 0 0 3 3m0-4a1 1 0 1 1-1 1a1 1 0 0 1 1-1"/><path d="M26 4H6a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2m0 22H6v-6l5-5l5.59 5.59a2 2 0 0 0 2.82 0L21 19l5 5Zm0-4.83l-3.59-3.59a2 2 0 0 0-2.82 0L18 19.17l-5.59-5.59a2 2 0 0 0-2.82 0L6 17.17V6h20Z"/></svg>`,
    'text': `<svg xmlns="http://www.w3.org/2000/svg" width="128" height="128" viewBox="0 0 24 24"><path d="M5 21q-.825 0-1.412-.587T3 19V5q0-.825.588-1.412T5 3h10l6 6v10q0 .825-.587 1.413T19 21zm2-4h10v-2H7zm0-4h10v-2H7zm0-4h7V7H7z"/></svg>`,
    'unknown': `<svg xmlns="http://www.w3.org/2000/svg" width="128" height="128" viewBox="0 0 24 24"><path d="M17 22q-2.075 0-3.537-1.463T12 17t1.463-3.537T17 12t3.538 1.463T22 17t-1.463 3.538T17 22m0-2q.275 0 .463-.187t.187-.463t-.187-.462T17 18.7t-.462.188t-.188.462t.188.463T17 20m-.45-1.9h.9v-.25q0-.275.15-.488t.35-.412q.35-.3.55-.575t.2-.775q0-.725-.475-1.162T17 14q-.575 0-1.037.338t-.663.912l.8.35q.075-.3.313-.525T17 14.85q.375 0 .588.188t.212.562q0 .275-.15.463t-.35.387q-.15.15-.312.3t-.288.35q-.075.15-.112.3t-.038.35zM12 9h5l-5-5l5 5l-5-5zM5 22q-.825 0-1.412-.587T3 20V4q0-.825.588-1.412T5 2h8l6 6v2.3q-.5-.15-1-.225T17 10q-1.425 0-2.687.538T12.1 12H7v2h3.675q-.225.475-.375.975T10.075 16H7v2h3.075q.175 1.125.7 2.163T12.125 22z"/></svg>`
};

window.fileTypeOrder = ['dir', 'package', 'video', 'audio', 'image', 'text', 'unknown'];

window.fileTypeOrderMap = {
    'dir': 0,
    'package': 1,
    'video': 2,
    'audio': 3,
    'image': 4,
    'text': 5,
    'unknown': 6
}

window.ajax = {
    getJson: function (url, callback) {
        const xhr = new XMLHttpRequest();
        xhr.open('get', url, true);
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.addEventListener('readystatechange', (e) => {
            if (xhr.readyState === 4 && xhr.status === 200) {
                callback && callback(JSON.parse(xhr.responseText));
            }
        });
        xhr.send();
    },

    postJson: function (url, data, callback) {
        const xhr = new XMLHttpRequest();
        xhr.open('post', url, true);
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.addEventListener('readystatechange', (e) => {
            if (xhr.readyState === 4 && xhr.status === 200) {
                callback && callback(JSON.parse(xhr.responseText));
            }
        });
        xhr.send(JSON.stringify(data));
    }
}

function htmlToElement(html) {
    const fatherDiv = document.createElement('div');
    fatherDiv.innerHTML = html;
    return fatherDiv.children[0];
}

function parseSearchParams() {
    let search = location.search;
    if (search === '' || search === '?') {
        return {}
    }
    search = search.slice(1);

    const params = search.split('&');
    const result = {};
    for (let i = 0; i < params.length; i++) {
        let param = params[i];
        let index = param.indexOf('=');
        let key = param.slice(0, index);
        let value = param.slice(index + 1);
        result[key] = value;
    }

    return result;
}

function connectPath(path1, path2) {
    if (path1.endsWith('/')) {
        path1 = path1.slice(0, path1.length - 1);
    }
    if (path2.startsWith('/')) {
        path2 = path2.slice(1, path2.length);
    }
    return path1 + '/' + path2;
}

function getFileTypeFromExtension(filename) {
    const i = filename.lastIndexOf('.');
    if (i === -1) {
        return 'unknown';
    }
    const extensionName = filename.slice(i + 1);
    const fileType = extensionTypeMap[extensionName];

    return fileType === undefined ? 'unknown' : fileType;
}

function getHumanFileSize(size) {
    const unitArray = ['B', 'KB', 'MB', 'GB', 'TB'];
    let i = 0;
    while (size >= 1024) {
        size = size / 1024;
        i = i + 1;
    }
    return size.toFixed(2) + ' ' + unitArray[i];
}

function getUpLevelPath(abspath) {
    abspath = abspath.replace(/\\/g, '/');
    if (abspath.endsWith('/')) {
        abspath = abspath.slice(0, abspath.length - 1);
    }
    const i = abspath.lastIndexOf('/');

    return i <= 0 ? abspath : abspath.slice(0, i);
}

function inArray(value, array) {
    for (let i = 0; i < array.length; i++) {
        if (array[i] === value) {
            return true;
        }
    }
    return false;
}