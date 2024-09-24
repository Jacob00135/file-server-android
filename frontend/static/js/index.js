(() => {

    renderContent();

    function renderContent() {
        ajax.getJson('/api/index', (response) => {
            const files = response['files'];
    
            if (files.length === 0) {
                return undefined;
            }

            const fileListEle = document.getElementById('file_list');
            fileListEle.querySelector('.empty-hint-text').remove();
    
            for (let i = 0; i < files.length; i++) {
                let f = document.createElement('a');
                f.className = 'file';
                f.textContent = files[i];

                f.href = `/?visible_dir=${files[i]}`;

                fileListEle.appendChild(f);
            }
        });
    }
})();