module.exports = ['$http', 'Restangular',
  function($http, Restangular) {
    var service = Restangular.service('api/categories');

    return {
      EntityNames: {
        SingularPascalCase: 'Category',
        PluralPascalCase: 'Categories',
        SingularCamelCase: 'category',
        PluralCamelCase: 'categories',
        SingularSnakeCase: 'category',
        PluralSnakeCase: 'categories'
      },
      getAll: function(args) {
        return service.getList(args);
      },
      getOne: function(id) {
        return service.one(id).get();
      },
      post: function(category) {
        return service.post(category);
      },
      put: function(category) {
        return category.put();
      },
      delete: function(category) {
        return category.remove();
      }
    };
  }
];