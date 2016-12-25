$(document).ready(function() {

    var freeDrawingBtn = $("#freeDrawing");
    var premadeShapesBtn = $("#premadeShapes");
    var lineWidthSlider = $("#lineWidthSlider");
    var colorPicker = $("#cp2");
    var clearCanvas = $("#clearCanvas");
    var undoItem = $("#undoItem");

    /*Line Width slider initialization*/
    var slider = lineWidthSlider.slider({
        formatter: function(value) {
            return "Current value: " + value;
        }
    });

    var canvas = new fabric.Canvas('myCanvas', {
        isDrawingMode: false
    });

    freeDrawingBtn.on("click", function () {
       if(premadeShapesBtn.hasClass("active")) {
           premadeShapesBtn.removeClass("active");
           premadeShapesBtn.html("Premade Shapes");
       }
       canvas.isDrawingMode = !canvas.isDrawingMode;
       $(this).toggleClass("active");
       if ($(this).hasClass("active")) {
           freeDrawingBtn.html("Cancel Drawing Mode");
       } else {
           freeDrawingBtn.html("Free Drawing Mode");
       }

    });

    premadeShapesBtn.on("click", function () {
        if(freeDrawingBtn.hasClass("active")) {
            freeDrawingBtn.removeClass("active");
            freeDrawingBtn.html("Free Drawing Mode");
        }
        canvas.isDrawingMode = false;
        $(this).toggleClass("active");
        if ($(this).hasClass("active")) {
            premadeShapesBtn.html("Cancel Drawing Mode");
        } else {
            premadeShapesBtn.html("Premade Shapes");
        }
    });

    /*Canvas Brushes*/
    var vLinePatternBrush = new fabric.PatternBrush(canvas);
    vLinePatternBrush.getPatternSrc = function () {

        var patternCanvas = fabric.document.createElement("canvas");
        patternCanvas.width = patternCanvas.height = 10;
        var ctx = patternCanvas.getContext("2d");

        ctx.strokeStyle = this.color;
        ctx.lineWidth = 5;
        ctx.beginPath();
        ctx.moveTo(0, 5);
        ctx.lineTo(10, 5);
        ctx.closePath();
        ctx.stroke();

        return patternCanvas;
    };

    var hLinePatternBrush = new fabric.PatternBrush(canvas);
    hLinePatternBrush.getPatternSrc = function () {

        var patternCanvas = fabric.document.createElement("canvas");
        patternCanvas.width = patternCanvas.height = 10;
        var ctx = patternCanvas.getContext("2d");

        ctx.strokeStyle = this.color;
        ctx.lineWidth = 5;
        ctx.beginPath();
        ctx.moveTo(5, 0);
        ctx.lineTo(5, 10);
        ctx.closePath();
        ctx.stroke();

        return patternCanvas;
    };

    var squarePatternBrush = new fabric.PatternBrush(canvas);
    squarePatternBrush.getPatternSrc = function() {

        var squareWidth = 10, squareDistance = 2;

        var patternCanvas = fabric.document.createElement('canvas');
        patternCanvas.width = patternCanvas.height = squareWidth + squareDistance;
        var ctx = patternCanvas.getContext('2d');

        ctx.fillStyle = this.color;
        ctx.fillRect(0, 0, squareWidth, squareWidth);

        return patternCanvas;
    };

    $("input[name=optradio]").on("change", function() {
        if(this.value === "hline") {
            canvas.freeDrawingBrush = vLinePatternBrush;
        } else if (this.value === "vline") {
            canvas.freeDrawingBrush = hLinePatternBrush;
        } else if (this.value === "sqline") {
            canvas.freeDrawingBrush = squarePatternBrush;
        } else {
            canvas.freeDrawingBrush = new fabric[this.value + "Brush"](canvas);
        }

        setCanvasParameters();
    });

    lineWidthSlider.on("slideStop", function () {
        setCanvasParameters();
    });

    colorPicker.on("changeColor", function () {
       setCanvasParameters();
    });

    var setCanvasParameters = (function setCanvasParameters() {
        if (canvas.freeDrawingBrush) {
            canvas.freeDrawingBrush.color = $("#cp2").colorpicker("getValue", "#000000");
            canvas.freeDrawingBrush.width = parseInt(slider[0].value, 10) || 1;
        }
        return setCanvasParameters;
    })();

    clearCanvas.on("click", function () {
       canvas.clear();
    });

    //Delete selected object with pressing the delete key
    $("html").keyup(function(e) {
        if(e.keyCode == 46 && canvas.getActiveObject()) {
            canvas.getActiveObject().remove();
        }
    });

    //Undo button
    undoItem.on("click", function () {
        var item = canvas.item(canvas.getObjects().length-1);
        canvas.remove(item);
    });

});
