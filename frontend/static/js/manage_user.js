(() => {

    main();

    function main() {
        // region 测试用代码
        const data = {
            'users': [
                {'username': 'admin', 'permission': 4},
                {'username': 'red', 'permission': 2},
                {'username': 'green', 'permission': 2},
                {'username': 'blue', 'permission': 2}
            ]
        };
        renderContent(data);
        // endregion
    }

    function renderContent(data) {
        const mainElement = document.getElementById('main');
        const users = data['users'];
        for (let i = 0; i < users.length; i++) {
            let user = users[i];
            let html = `<div class="row">
                <div class="item">${user['username']}</div>
                <div class="item">${permissionMap[user['permission']]}</div>
            </div>`;

            mainElement.appendChild(htmlToElement(html));
        }
    }
})();