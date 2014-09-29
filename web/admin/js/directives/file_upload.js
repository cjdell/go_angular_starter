module.exports = ['$q', 'Uploader', 'Utility',
  function($q, Uploader, Utility) {
    return {
      scope: {
        fileUpload: "=",
        fileUploadIndicator: "="
      },
      template: '<input type="file" class="ng-hide" /><a href="#" class="pure-button select-image-button">Select Image</a>',
      link: function(scope, element, attributes) {
        var inputFile = element.find('input');
        var button = element.find('a');

        element.addClass('file-upload');

        var deferred = $q.defer();

        deferred.resolve(null); // Default promise is that no file has been choosen

        scope.fileUpload = deferred.promise;

        button.bind('click', function(e) {
          e.preventDefault();

          inputFile[0].click();
        });

        inputFile.bind('change', function(changeEvent) {
          if (changeEvent.target.files.length === 0) return;

          scope.fileUploadIndicator = "Uploading";

          var fileId = Utility.guid();

          var promise = Uploader.uploadBlob(changeEvent.target.files[0], fileId).then(function(savedFileName) {
            scope.fileUploadIndicator = "Upload Complete: " + savedFileName;

            return savedFileName;
          });

          scope.$apply(function() {
            scope.fileUpload = promise; // Write the upload promise to the scope
          });
        });
      }
    };
  }
];