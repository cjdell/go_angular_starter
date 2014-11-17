module.exports = ['$http', 'Restangular', function($http, Restangular) {
    var service = Restangular.service('api/<%= entityNamePluralSnakeCase %>');

    return {
      EntityNames: {
        SingularPascalCase: '<%= entityNameSingularPascalCase %>',
        PluralPascalCase: '<%= entityNamePluralPascalCase %>',
        SingularCamelCase: '<%= entityNameSingularCamelCase %>',
        PluralCamelCase: '<%= entityNamePluralCamelCase %>',
        SingularSnakeCase: '<%= entityNameSingularSnakeCase %>',
        PluralSnakeCase: '<%= entityNamePluralSnakeCase %>'
      },
      getAll: function(args) {
        return service.getList();
      },
      getOne: function(id) {
        return service.one(id).get();
      },
      post: function(<%= entityNameSingularCamelCase %>) {
        return service.post(<%= entityNameSingularCamelCase %>);
      },
      put: function(<%= entityNameSingularCamelCase %>) {
        return <%= entityNameSingularCamelCase %>.put();
      },
      delete: function(<%= entityNameSingularCamelCase %>) {
        return <%= entityNameSingularCamelCase %>.remove();
      }
    };
  }
];