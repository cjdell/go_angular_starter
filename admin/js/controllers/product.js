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

    ctrl.paramsForLoad = {
      owned_only: 1
    };

    // Bootstrap the controller
    ctrl.init();
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

    $scope.$on('item-saving', (e, record, promises) => {
      // record.Changes.AttributeStore = getAttributeStore();
      record.Changes.ImageChanges = getImageChanges();

      // Wait for new image if still uploading...
      promises.push($scope.newImage().then((savedFileName) => {
        record.Changes.NewImageHandle = $scope.newImageHandle;
        record.Changes.NewImageFileName = savedFileName;
      }));
    });

    $scope.$watchCollection('record.Categories', () => {
      if ($scope.record && $scope.record.Categories) {
        $scope.record.CategoryIds = Utility.toPgArray(_.map($scope.record.Categories, (c) => c.Id));
      }
    });

    function getImageChanges() {
      var imageChanges = {};

      _.each($scope.record.Images, function(image) {
        imageChanges[image.Handle] = {
          Desc: image.Desc
        };
      });

      return imageChanges;
    }

    // Bootstrap the controller
    ctrl.init();
  }
];

module.exports = {
  ProductsController: ProductsController,
  ProductController: ProductController
};