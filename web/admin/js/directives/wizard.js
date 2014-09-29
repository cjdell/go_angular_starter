module.exports = ['$state', '$stateParams',
  function($state, $stateParams) {
    return {
      restrict: 'C',
      link: function(scope, element, attrs) {
        scope.goToStage = function(toStageIndex) {
          // Add the class to perform the correct CSS animations
          if ($stateParams.stageIndex < toStageIndex) {
            element.addClass('wizard-moved-forward');
            element.removeClass('wizard-moved-backward');
          } else if ($stateParams.stageIndex > toStageIndex) {
            element.addClass('wizard-moved-backward');
            element.removeClass('wizard-moved-forward');
          }

          $state.go($state.current.name, {
            stageIndex: toStageIndex
          });
        };
      },
    };
  }
];