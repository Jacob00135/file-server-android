(() => {

    main();

    function main() {
        document.querySelector('#login-widget form').addEventListener('submit', loginFormSubmitEvent);
    }

    function loginFormSubmitEvent(e) {
        e.preventDefault();

        const form = document.querySelector('#login-widget form');
        const data = {
            'username': form['username'].value,
            'password': form['password'].value
        };
        ajax.postJson(form.action, data, (response) => {
            if (response['success']) {
                location.assign('/');
                return undefined;
            }

            const errorMessageElement = document.getElementById('error-message');
            errorMessageElement.classList.remove('hidden');
            errorMessageElement.textContent = response['message'];
        });
    }

})();