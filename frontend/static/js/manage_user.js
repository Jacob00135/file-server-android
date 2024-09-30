(() => {

    main();

    function main() {
        document.getElementById('add-user').addEventListener('click', openAddUserDialog);
        document.querySelector('#add-user-dialog form').addEventListener('submit', addUserFormSubmitEvent);
        document.querySelector('#add-user-dialog form .btn-group .close').addEventListener('click', closeAddUserDialog);

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
            if (user['permission'] === 4) {
                continue;
            }

            let html = `<div class="row" data-user-id="${user['id']}">
                <div class="item username">${user['username']}</div>
                <div class="item permission">
                    <div class="text">${permissionMap[user['permission']]}</div>
                </div>
                <button class="item delete-btn" type="button">
                    <svg class="icon" xmlns="http://www.w3.org/2000/svg" width="128" height="128" viewBox="0 0 24 24"><path d="M19 4h-3.5l-1-1h-5l-1 1H5v2h14M6 19a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7H6z"/></svg>
                </button>
            </div>`;

            let row = htmlToElement(html);
            row.querySelector('.delete-btn').addEventListener('click', deleteButtonClickEvent);

            mainElement.appendChild(row);
        }
    }

    function deleteButtonClickEvent(e) {
        if (!confirm('确定删除此记录？')) {
            return undefined;
        }

        let ele = e.target;
        let userId = ele.getAttribute('data-user-id');
        while (userId === null) {
            ele = ele.parentNode;
            userId = ele.getAttribute('data-user-id');
        }

        ajax.deleteRequest(`/api/manage_user/${userId}`, (xhr) => {
            const response = JSON.parse(xhr.responseText);
            if (!response['success']) {
                alert(`删除失败：${response['message']}`);
                return undefined;
            }

            location.reload();
        });
    }

    function openAddUserDialog() {
        document.getElementById('overlay').classList.remove('hidden');
        document.getElementById('add-user-dialog').classList.remove('hidden');
    }

    function closeAddUserDialog() {
        document.getElementById('overlay').classList.add('hidden');
        document.getElementById('add-user-dialog').classList.add('hidden');
    }

    function addUserFormSubmitEvent(e) {
        e.preventDefault();

        const form = e.target;
        const errorMessageElement = form.querySelector('#add-user-error-message');
        const password = form.querySelector('[name="password"]').value;
        const againPassword = form.querySelector('#again-password').value;
        if (password !== againPassword) {
            errorMessageElement.textContent = '两次输入的密码不一致！';
            errorMessageElement.classList.remove('hidden');
            return undefined;
        }

        const url = form.action;
        const data = {
            'username': form.querySelector('[name="username"]').value,
            'password': password
        };
        ajax.postJson(url, data, (response) => {
            if (response['success']) {
                location.reload();
            } else {
                errorMessageElement.textContent = response['message'];
                errorMessageElement.classList.remove('hidden');
            }
        });
    }
})();