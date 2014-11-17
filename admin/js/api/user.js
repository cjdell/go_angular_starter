module.exports = ['$http', 'Restangular',
  function($http, Restangular) {
    var service = Restangular.service('api/users');

    return {
      EntityNames: {
        SingularPascalCase: 'User',
        PluralPascalCase: 'Users',
        SingularCamelCase: 'user',
        PluralCamelCase: 'users',
        SingularSnakeCase: 'user',
        PluralSnakeCase: 'users'
      },
      getAll: function(args) {
        return service.getList();
      },
      getOne: function(id) {
        return service.one(id).get();
      },
      post: function(user) {
        return service.post(user);
      },
      put: function(user) {
        return user.put();
      },
      delete: function(user) {
        return user.remove();
      }
    };
  }
];