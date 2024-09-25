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