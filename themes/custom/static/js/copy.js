$(function() {
    var selector = 'div.markdown-modules5 code';
    // go through all code blocks in the 
    // markdown and add buttons to copy
    $(selector).each(function(idx, val) {
        var valID = `code-${idx}`;
        $(this).attr("id", valID);

        var newBtn = document.createElement("button");
        newBtn.className = "btn btn-link btn-sm";
        newBtn.innerHTML = "copy";
        newBtn.setAttribute("data-clipboard-target", `#${valID}`);
        // val.after(newBtn);
    });

    // find all code blocks in the markdown
    var clipboard = new ClipboardJS(selector);

    clipboard.on('success', function(e) {
        console.info('Action:', e.action);
        console.info('Text:', e.text);
        console.info('Trigger:', e.trigger);
    });
    
    clipboard.on('error', function(e) {
        console.error('Action:', e.action);
        console.error('Trigger:', e.trigger);
    });
    
})