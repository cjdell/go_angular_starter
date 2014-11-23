// Floats an element when it is about to scroll out of view
module.exports = [

  function() {
    return {
      restrict: 'C',
      link: function(scope, element, attrs) {
        var top = 0;

        function getTop() {
          element.removeClass('fixed');
          element.parent().removeClass('has-floated');
          top = getOffsetTop(element[0]);
        }

        function checkScroll() {
          var fixed = window.pageYOffset > top;
          element.toggleClass('fixed', fixed);
          element.parent().toggleClass('has-floated', fixed);
        }

        setTimeout(function() {
          getTop();
          checkScroll();
        }, 100);

        window.addEventListener('scroll', checkScroll);

        window.addEventListener('resize', getTop);
      },
    };
  }
];

function getOffsetTop(elem) {
  var offsetTop = 0;
  do {
    if (!isNaN(elem.offsetTop)) {
      offsetTop += elem.offsetTop;
    }
    elem = elem.offsetParent;
  } while (elem);
  return offsetTop;
}