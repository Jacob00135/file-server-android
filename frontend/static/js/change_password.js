(() => {

    main();

    function main() {
        document.querySelector('#change-password-widget form').addEventListener('submit', changePasswordFormSubmitEvent);
    }

    function changePasswordFormSubmitEvent(e) {
        e.preventDefault();

        const form = document.querySelector('#change-password-widget form');
        const errorMessageElement = document.getElementById('error-message');
        const password = form.querySelector('[name="password"]').value;
        const againPassword = form.querySelector('#again-password').value;
        if (password !== againPassword) {
            errorMessageElement.textContent = '两次输入的密码不一致！';
            errorMessageElement.classList.remove('hidden');
            return undefined;
        }

        const data = {'password': password};
        ajax.postJson(form.action, data, (response) => {
            if (response['success']) {
                location.assign('/login');
                return undefined;
            }

            errorMessageElement.textContent = response['message'];
            errorMessageElement.classList.remove('hidden');
        });
    }

})();