var assetData = $("#XXXXX").formobj();　　
for (var key in assetData) {　　
    var lowerkey = key.toLowerCase();　　 //赋给新的属性名，删除旧的
    　　
    assetData[lowerkey] = assetData[key];　　
    delete assetData[key];
}
console.log(assetData)