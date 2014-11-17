module.exports = ['$http', 'Restangular',
  function($http, Restangular) {
    var service = Restangular.service('api/products');

    return {
      EntityNames: {
        SingularPascalCase: 'Product',
        PluralPascalCase: 'Products',
        SingularCamelCase: 'product',
        PluralCamelCase: 'products',
        SingularSnakeCase: 'product',
        PluralSnakeCase: 'products'
      },
      getAll: function(args) {
        return service.getList(args);
      },
      getOne: function(id) {
        return service.one(id).get();
      },
      post: function(product) {
        return service.post(product);
      },
      put: function(product) {
        return product.put();
      },
      delete: function(product) {
        return product.remove();
      }
    };
  }
];