var Validate = require('validate-arguments');

module.exports = ['$http',
  function($http) {
    return {
      EntityNames: {
        SingularPascalCase: 'Category',
        PluralPascalCase: 'Categories',
        SingularCamelCase: 'category',
        PluralCamelCase: 'categories',
        SingularSnakeCase: 'category',
        PluralSnakeCase: 'categories'
      },
      GetAll: function() {
        return $http.jsonrpc('/api', 'CategoryApi.GetAll', [{}]);
      },
      GetChildren: function(parentId) {
        var args = Validate.validate(arguments, ['natural']);

        if (!args.isValid()) {
          return $q.reject(args.errorString());
        }

        return $http.jsonrpc('/api', 'CategoryApi.GetChildren', [{
          ParentId: parentId
        }]);
      },
      GetOne: function(id) {
        var args = Validate.validate(arguments, ['natural']);

        if (!args.isValid()) {
          return $q.reject(args.errorString());
        }

        return $http.jsonrpc('/api', 'CategoryApi.GetOne', [{
          Id: id
        }]);
      },
      Insert: function(category) {
        return $http.jsonrpc('/api', 'CategoryApi.Insert', [{
          Category: category
        }]);
      },
      Update: function(category) {
        return $http.jsonrpc('/api', 'CategoryApi.Update', [{
          Category: category
        }]);
      },
      Delete: function(id) {
        return $http.jsonrpc('/api', 'CategoryApi.Delete', [{
          Id: id
        }]);
      }
    };
  }
];