/*
引入模块 必须遵循AMD规范
*/

require.config({
    paths: {
        father: "../js/amd_father",
        chilF: "../js/amd_child",
        domReady: "../pkgs/domReady"
    }
})

require(["domReady!", "father", "child"], function (doc, father, child) {
    var ret = father.Add(10, 1);
    console.log(ret);
    father.Show();
    console.log(child.Double(2, 3, 4));
})