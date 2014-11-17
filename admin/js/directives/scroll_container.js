// Stops overflow divs from scrolling their parents
module.exports = [

  function() {
    return {
      restrict: 'C',
      link: function(scope, element, attrs) {
        element.bind('mousewheel', function(e) {
          var d = e.wheelDeltaY;

          if (window.innerWidth < 768) return;

          if ((this.scrollTop === (this.scrollHeight - this.clientHeight) && d < 0) || (this.scrollTop === 0 && d > 0)) {
            e.preventDefault();
          }
        });
      },
    };
  }
];