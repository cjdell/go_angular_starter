var _ = require('underscore'),
  Validate = require('validate-arguments');

var ListController = ['$scope', '$q', 'Api',
  function($scope, $q, Api) {
    var args = Validate.validate(arguments, ['object', 'object', {
      EntityNames: 'object',
      getAll: 'function'
    }]);

    if (!args.isValid()) {
      throw args.errorString();
    }

    var ctrl = this; // Only time `this` is used. Save us from confusion...

    ctrl.paramsForLoad = {}; // Maps to query string for API GET request

    ctrl.init = function() {
      console.log('ListController: init');
      // Load the list
      ctrl.load();
    };

    ctrl.load = function() {
      Api.getAll(ctrl.paramsForLoad).then(ctrl.loaded, ctrl.failed);
    };

    ctrl.loaded = function(reply) {
      $scope.records = reply;

      // console.table($scope.records);

      $scope.$emit('itemCountDiscovered', Api.EntityNames.SingularPascalCase, $scope.records.length);
    };

    ctrl.failed = function(err) {
      console.error(err);
    };

    $scope.$on('item-saved', () => ctrl.load());

    $scope.$on('item-deleted', () => ctrl.load());
  }
];

var ItemController = ['$scope', '$q', '$state', '$stateParams', 'Api',
  function($scope, $q, $state, $stateParams, Api) {
    var args = Validate.validate(arguments, ['object', 'object', 'object', 'object', {
      EntityNames: 'object',
      getOne: 'function'
    }]);

    if (!args.isValid()) {
      throw args.errorString();
    }

    var ctrl = this; // Only time `this` is used. Save us from confusion...

    ctrl.init = function() {
      console.log('ItemController: init');
      // If there is an ID, load the record from the API, otherwise, start with a new instance
      if ($stateParams.id !== undefined) {
        ctrl.load(parseInt($stateParams.id));
      } else {
        ctrl.blank();
      }
    };

    ctrl.load = function(id) {
      Api.getOne(id).then(ctrl.loaded, ctrl.failed);
    };

    ctrl.blank = function() {
      $scope.record = {};

      beginObserve($scope.record);
    };

    ctrl.new = function() {
      ctrl.goToRecord();
    };

    ctrl.validate = function() {
      return true;
    };

    ctrl.save = function() {
      if (ctrl.validate()) {
        $scope.$emitp('item-saving', $scope.record).then(function(results) {
          $scope.saveClass = 'saving';

          if ($scope.record.Id !== undefined) {
            ctrl.update();
          } else {
            ctrl.insert();
          }
        });
      }
    };

    ctrl.update = function() {
      $scope.$emitp('item-updating', $scope.record)
        .then(function(results) {
          return Api.put($scope.record);
        })
        .then(ctrl.saved, ctrl.failed);
    };

    ctrl.insert = function() {
      $scope.$emitp('item-inserting', $scope.record)
        .then(function(results) {
          return Api.post($scope.record);
        })
        .then(ctrl.saved, ctrl.failed);
    };

    ctrl.delete = function() {
      $scope.$emitp('item-deleting', $scope.record)
        .then(function(results) {
          return Api.delete($scope.record);
        })
        .then(ctrl.deleted, ctrl.failed);
    };

    // Callbacks are declared privately
    ctrl.loaded = function(reply) {
      $scope.record = reply;

      beginObserve($scope.record);

      console.log('Loaded Record:', $scope.record);
    };

    ctrl.failed = function(err) {
      if (err && err.data) alert(err.data.Error);
      console.error(err);
    };

    ctrl.saved = function(reply) {
      $scope.record = reply;

      $scope.saveClass = 'saved';

      beginObserve($scope.record);

      console.log('Saved Record:', $scope.record);

      // Broadcast saved event
      $scope.$emit('item-saved', $scope.record);

      ctrl.goToRecord($scope.record.Id);
    };

    ctrl.deleted = function() {
      console.log('Deleted Record:', $scope.record);

      // Broadcast deleted event
      $scope.$emit('item-deleted', $scope.record);

      // Go to a blank record
      ctrl.goToRecord();
    };

    ctrl.goToRecord = function(id) {
      if (id) {
        $state.go('records.' + Api.EntityNames.PluralSnakeCase + '.view', {
          id: id
        });
      } else {
        $state.go('records.' + Api.EntityNames.PluralSnakeCase + '.new');
      }
    };

    // These are the methods we choose to expose to the view via the $scope
    $scope.new = () => ctrl.new();
    $scope.save = () => ctrl.save();
    $scope.delete = () => ctrl.delete();

    function beginObserve(record) {
      record.Changes = {
        Fields: []
      };

      var observer = new ObjectObserver(record);

      observer.open(function(added, removed, changed, getOldValueFn) {
        Object.keys(added).forEach(function(property) {
          record.Changes.Fields.push(property);
        });

        Object.keys(changed).forEach(function(property) {
          record.Changes.Fields.push(property);
        });
      });

      // Object.observe(record, function(changes) {
      //   record.Changes.Fields.push(changes[0].name);
      // });
    }
  }
];

module.exports = {
  ListController: ListController,
  ItemController: ItemController
};