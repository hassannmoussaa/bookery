require(["jquery","requestjs","alertjs","moment","dot","urlManager","popupjs","cp","app/modules/moment-ar","imgbox","domurl","dropdown","loader","common"],function(e,a,s,p,t,r,o){s.init("#alert","ltr"),e.ajaxSetup({headers:{"X-Csrf-Token":_CSRFToken}}),t.templateSettings={evaluate:/<%([\s\S]+?)%>/g,interpolate:/<%=([\s\S]+?)%>/g,encode:/<%!([\s\S]+?)%>/g,use:/<%#([\s\S]+?)%>/g,define:/<%##\s*([\w\.$]+)\s*(\:|=)([\s\S]+?)#%>/g,conditional:/<%\?(\?)?\s*([\s\S]*?)\s*%>/g,iterate:/<%~\s*(?:%>|([\s\S]+?)\s*\:\s*([\w$]+)\s*(?:\:\s*([\w$]+))?\s*%>)/,varname:"it",strip:!0,append:!0,selfcontained:!1},o.init("#popup"),_IsCPPage?require(["app/cp-pages/"+_PageName.replace("cp-","")]):require(["app/pages/"+_PageName,"app/modules/date"])});