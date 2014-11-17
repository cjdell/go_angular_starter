// Locks an element to its parent's width
module.exports = [

  function() {
    return {
      restrict: 'C',
      link: function(scope, element, attrs) {
        resize();
        setInterval(resize, 100);

        function resize() {
          var width = element[0].parentNode.clientWidth;
          element[0].style.width = width + 'px';
        }
      },
    };
  }
];