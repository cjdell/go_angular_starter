module.exports = ['$q', 'Uploader', 'Utility',
  function($q, Uploader, Utility) {
    function noFile() {
      var deferred = $q.defer();
      deferred.resolve(null);
      return deferred.promise;
    }

    return {
      scope: {
        fileUpload: "="
      },
      templateUrl: 'views/directives/file_upload.html',
      link: function(scope, element, attributes) {
        var inputFile = element.find('input');
        var button = element.find('a');

        element.addClass('file-upload');

        // Default promise is that no file has been choosen
        scope.fileUpload = noFile;

        scope.uploading = false;

        button.bind('click', function(e) {
          e.preventDefault();

          inputFile[0].click();
        });

        inputFile.bind('change', function(changeEvent) {
          if (changeEvent.target.files.length === 0) return;

          scope.fileUploadIndicator = "Uploading...";
          scope.uploading = true;

          var fileId = Utility.guid();

          var promise = Uploader.uploadBlob(changeEvent.target.files[0], fileId).then(function(savedFileName) {
            scope.fileUploadIndicator = "Upload Complete";
            scope.uploading = false;

            return savedFileName;
          });

          scope.$apply(function() {
            // Write the upload promise to the scope
            scope.fileUpload = function() {
              scope.fileUpload = noFile;
              return promise;
            };
          });
        });
      }
    };
  }
];