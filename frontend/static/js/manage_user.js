(() => {

    main();

    function main() {
        ajax.getJson('/api/manage_user', (response) => {
            renderContent(response['users']);
        });

        // region 测试用代码
        /*const data = {
            'users': [
                {'username': 'admin', 'permission': 4},
                {'username': 'red', 'permission': 2},
                {'username': 'green', 'permission': 2},
                {'username': 'blue', 'permission': 2}
            ]
        };
        renderContent(data);*/
        // endregion
    }

    function renderContent(users) {
        const mainElement = document.getElementById('main');
        for (let i = 0; i < users.length; i++) {
            let user = users[i];
            let html = `<div class="row">
                <div class="item username">${user['username']}</div>
                <div class="item permission">
                    <div class="text">${permissionMap[user['permission']]}</div>
                </div>
                <button class="item delete-btn" type="button">
                    <svg class="icon" xmlns="http://www.w3.org/2000/svg" width="128" height="128" viewBox="0 0 24 24"><path d="M19 4h-3.5l-1-1h-5l-1 1H5v2h14M6 19a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7H6z"/></svg>
                </button>
            </div>`;

            mainElement.appendChild(htmlToElement(html));
        }
    }
})();