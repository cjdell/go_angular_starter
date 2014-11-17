module.exports = ['RestangularProvider',
  function(RestangularProvider) {
    // Records have Id field (capital I)
    RestangularProvider.setRestangularFields({
      id: 'Id'
    });

    // Remove request body for DELETE
    RestangularProvider.setRequestInterceptor(function(elem, operation) {
      if (operation === "remove") {
        return null;
      }
      return elem;
    });
  }
];