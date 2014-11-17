module.exports = ['$provide',
  function($provide) {
    $provide.decorator('$http', ['$delegate', '$q',
      function($delegate, $q) {
        $delegate.jsonrpc = function(url, method, parameters, config) {
          var deferred = $q.defer();

          var data = {
            "jsonrpc": "2.0",
            "method": method,
            "params": parameters,
            "id": 1
          };

          $delegate.post(url, data, angular.extend({
            'headers': {
              'Content-Type': 'application/json'
            }
          }, config)).then(succeeded, failed);

          function succeeded(data) {
            if (!data.data.error) {
              deferred.resolve(data.data.result);
            } else {
              console.error(data.data.error);
              deferred.reject(data.data.error);
            }
          }

          function failed(error) {
            if (typeof error.data === 'string') {
              alert(error.data); // Useful for debugging
            }

            console.error(error);
            deferred.reject(error);
          }

          return deferred.promise;
        };

        return $delegate;
      }
    ]);
  }
];