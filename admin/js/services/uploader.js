module.exports = ['$q',
  function($q) {
    function uploadBlob(blob, fileId) {
      var deferred = $q.defer();

      if (!(blob instanceof Blob)) {
        console.error("Not a blob!");
        return;
      }

      if (typeof fileId !== 'string') {
        console.error("No file ID!");
        return;
      }

      var xhr = new XMLHttpRequest();

      xhr.onload = ajaxSuccess;
      xhr.onerror = ajaxError;

      xhr.open("post", "/upload", true);

      xhr.setRequestHeader("Content-Type", "application/octet-stream");

      if (blob.name) xhr.setRequestHeader("X-Upload-File-Name", blob.name);
      if (fileId) xhr.setRequestHeader("X-Upload-File-ID", fileId);

      xhr.send(blob);

      return deferred.promise;

      function ajaxSuccess() {
        deferred.resolve(xhr.responseText);
      }

      function ajaxError(error) {
        deferred.reject(error);
      }
    }

    return {
      uploadBlob: uploadBlob
    };
  }
];