var Validate = require('validate-arguments');

module.exports = ['$http',
  function($http) {
    return {
      EntityNames: {
        SingularPascalCase: 'User',
        PluralPascalCase: 'Users',
        SingularCamelCase: 'user',
        PluralCamelCase: 'users',
        SingularSnakeCase: 'user',
        PluralSnakeCase: 'users'
      },
      GetAll: function() {
        return $http.jsonrpc('/api', 'UserApi.GetAll', [{}]);
      },
      GetOne: function(id) {
        var args = Validate.validate(arguments, ['natural']);

        if (!args.isValid()) {
          return $q.reject(args.errorString());
        }

        return $http.jsonrpc('/api', 'UserApi.GetOne', [{
          Id: id
        }]);
      },
      Insert: function(user) {
        return $http.jsonrpc('/api', 'UserApi.Insert', [{
          User: user
        }]);
      },
      Update: function(user) {
        return $http.jsonrpc('/api', 'UserApi.Update', [{
          User: user
        }]);
      },
      Delete: function(id) {
        return $http.jsonrpc('/api', 'UserApi.Delete', [{
          Id: id
        }]);
      }
    };
  }
];