altStyle=window.location.hash.substr(1);
window.addEventListener("hashchange", changeStyle, false);
function changeStyle() {
	title = location.hash.substr(1);
	var stylesheet = document.getElementById("stylesheet");
	stylesheet.setAttribute("href", "/s/styles/"+title+".css");
} 
if(altStyle) changeStyle();
function getPastePre(){
	return document.getElementById('code');
}
function highlightPre(pastePre){
	hljs.highlightBlock(document.body);  	
	hljs.highlightBlock(pastePre);
	}
function initPaste(){
	inPastePre = getPastePre();
	highlightPre(inPastePre);
}
