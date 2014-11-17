module.exports = ['$rootScope', '$q',
  function($rootScope, $q) {
    $rootScope.$emitp = function(name, obj) {
      var promises = [];
      this.$emit(name, obj, promises);
      return $q.all(promises);
    };

    $rootScope.$broadcastp = function(name, obj) {
      var promises = [];
      this.$broadcast(name, obj, promises);
      return $q.all(promises);
    };
  }
];