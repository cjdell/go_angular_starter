var _ = require('underscore');

module.exports = [

  function() {
    function guid() {
      return s4() + s4() + '-' + s4() + '-' + s4() + '-' +
        s4() + '-' + s4() + s4() + s4();

      function s4() {
        return Math.floor((1 + Math.random()) * 0x10000)
          .toString(16)
          .substring(1);
      }
    }

    function fromPgArray(pgArray) {
      _.map(pgArray.substring(1, pgArray.length - 2).split(','), (id) => parseInt(id));
    }

    function toPgArray(array) {
      return '{' + array.join(',') + '}';
    }

    return {
      guid: guid,
      fromPgArray: fromPgArray,
      toPgArray: toPgArray
    };
  }
];