var angular = require('angular'),
  _ = require('underscore'),
  common = require('./common');

var ProductsController = ['$injector', '$scope', '$state', '$stateParams', 'ProductApi',
  function($injector, $scope, $state, $stateParams, ProductApi) {
    var ctrl = this;

    // Inherit shared functionality from the ListController
    $injector.invoke(common.ListController, ctrl, {
      $scope: $scope,
      Api: ProductApi
    });
  }
];

var ProductController = ['$injector', '$scope', '$state', '$stateParams', 'ProductApi', 'ngDialog', 'Utility',
  function($injector, $scope, $state, $stateParams, ProductApi, ngDialog, Utility) {
    var ctrl = this;

    // Inherit shared functionality from the ItemController
    $injector.invoke(common.ItemController, ctrl, {
      $scope: $scope,
      $state: $state,
      $stateParams: $stateParams,
      Api: ProductApi
    });

    // Override update
    ctrl.update = () => {
      // Wait for new image if still uploading...
      $scope.newImage.then((savedFileName) => {
        ProductApi.Update($scope.record, savedFileName).then(ctrl.saved, ctrl.failed);
      });
    };

    $scope.$watchCollection('record.Categories', () => {
      if ($scope.record && $scope.record.Categories) {
        $scope.record.CategoryIds = Utility.toPgArray(_.map($scope.record.Categories, (c) => c.Id));
      }
    });
  }
];

module.exports = {
  ProductsController: ProductsController,
  ProductController: ProductController
};