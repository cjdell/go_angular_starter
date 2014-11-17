module.exports = ['$rootScope',
  function($rootScope) {
    // Invoked when the count of a particular record type is discovered
    $rootScope.$on("itemCountDiscovered", function(e, type, count) {
      $rootScope.itemCounts = $rootScope.itemCounts || {};

      $rootScope.itemCounts[type] = count;
    });
  }
];