document.addEventListener('DOMContentLoaded', function() {
    const dropZone = document.getElementById('dropZone');
    const fileInput = document.getElementById('fileInput');
    const previewArea = document.getElementById('previewArea');
    const uploadForm = document.getElementById('uploadForm');
    const closeButton = document.getElementById('closeButton');
    const addMoreButton = document.getElementById('addMoreButton');

    // Handle click on drop zone to trigger file input
    dropZone.addEventListener('click', () => {
        fileInput.click();
    });

    // Handle file selection
    fileInput.addEventListener('change', handleFiles);

    // Handle drag over
    dropZone.addEventListener('dragover', (e) => {
        e.preventDefault();
        dropZone.classList.add('dragover');
    });

    // Handle drag leave
    dropZone.addEventListener('dragleave', () => {
        dropZone.classList.remove('dragover');
    });

    // Handle drop
    dropZone.addEventListener('drop', (e) => {
        e.preventDefault();
        dropZone.classList.remove('dragover');

        if (e.dataTransfer.files.length) {
            fileInput.files = e.dataTransfer.files;
            handleFiles({ target: fileInput });
        }
    });

    // Add more files
    addMoreButton.addEventListener('click', () => {
        fileInput.click();
    });

    // Close button
    closeButton.addEventListener('click', () => {
        window.location.href = '/'; // or wherever you want to redirect
    });

    // Handle file previews
    function handleFiles(e) {
        const files = e.target.files;
        previewArea.innerHTML = '';

        if (files.length > 0) {
            for (let i = 0; i < files.length; i++) {
                const file = files[i];
                const fileElement = document.createElement('div');
                fileElement.className = 'file-preview';

                fileElement.innerHTML = `
<div class="file-info">
    <span class="file-name">${file.name}</span>
    <span class="file-size">${formatFileSize(file.size)}</span>
</div>
<button class="remove-file" data-index="${i}">×</button>
`;

                previewArea.appendChild(fileElement);
            }

            // Add remove file functionality
            document.querySelectorAll('.remove-file').forEach(button => {
                button.addEventListener('click', function() {
                    const index = parseInt(this.getAttribute('data-index'));
                    removeFile(index);
                });
            });
        }
    }

    function removeFile(index) {
        const dt = new DataTransfer();
        const files = fileInput.files;

        for (let i = 0; i < files.length; i++) {
            if (i !== index) {
                dt.items.add(files[i]);
            }
        }

        fileInput.files = dt.files;
        handleFiles({ target: fileInput });
    }

    function formatFileSize(bytes) {
        if (bytes === 0) return '0 Bytes';
        const k = 1024;
        const sizes = ['Bytes', 'KB', 'MB', 'GB'];
        const i = Math.floor(Math.log(bytes) / Math.log(k));
        return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
    }
});

// Add this to your script section
uploadForm.addEventListener('submit', function(e) {
    e.preventDefault();
    
    const formData = new FormData(uploadForm);
    const files = fileInput.files;
    
    for (let i = 0; i < files.length; i++) {
        formData.append('document', files[i]);
    }
    
    const uploadStatus = document.getElementById('uploadStatus');
    uploadStatus.textContent = 'Uploading...';
    uploadStatus.style.display = 'block';
    
    fetch('/upload', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        if (data.message) {
            uploadStatus.textContent = data.message;
            uploadStatus.className = 'upload-status success';
            
            if (data.errors) {
                uploadStatus.textContent += '\nErrors: ' + data.errors.join(', ');
            }
            
            // Clear form after successful upload
            setTimeout(() => {
                uploadForm.reset();
                previewArea.innerHTML = '';
                uploadStatus.style.display = 'none';
            }, 3000);
        }
    })
    .catch(error => {
        uploadStatus.textContent = 'Upload failed: ' + error;
        uploadStatus.className = 'upload-status error';
    });
});
