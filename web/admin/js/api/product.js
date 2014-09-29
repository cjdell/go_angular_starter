var Validate = require('validate-arguments');

module.exports = ['$http', '$q',
  function($http, $q) {
    return {
      EntityNames: {
        SingularPascalCase: 'Product',
        PluralPascalCase: 'Products',
        SingularCamelCase: 'product',
        PluralCamelCase: 'products',
        SingularSnakeCase: 'product',
        PluralSnakeCase: 'products'
      },
      GetAll: function() {
        return $http.jsonrpc('/api', 'ProductApi.GetAll', [{}]);
      },
      GetOne: function(id) {
        var args = Validate.validate(arguments, ['natural']);

        if (!args.isValid()) {
          return $q.reject(args.errorString());
        }

        return $http.jsonrpc('/api', 'ProductApi.GetOne', [{
          Id: id
        }]);
      },
      Insert: function(product, newImageFileName) {
        return $http.jsonrpc('/api', 'ProductApi.Insert', [{
          Product: product,
          Extra: {
            NewImageFileName: newImageFileName
          }
        }]);
      },
      Update: function(product, newImageFileName) {
        return $http.jsonrpc('/api', 'ProductApi.Update', [{
          Product: product,
          Extra: {
            NewImageFileName: newImageFileName
          }
        }]);
      },
      Delete: function(id) {
        return $http.jsonrpc('/api', 'ProductApi.Delete', [{
          Id: id
        }]);
      }
    };
  }
];