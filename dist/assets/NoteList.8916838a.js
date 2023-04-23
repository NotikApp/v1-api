import{E as a,m as o,n as i,R as r,L as l,a as c,b as p}from"./index.56711988.js";const d={name:"NoteCard",components:{EditButton:a},props:["note"],data(){return{reveal:!1,modal:!1}},methods:{...o(["deleteNoteAction","editNoteAction"]),getBody(){return this.note.text.length>40?this.note.text.substring(0,39)+"...":this.note.text},deleteCurrent(){this.deleteNoteAction(this.note)},prompt(){this.$buefy.dialog.prompt({message:"Edit note",inputAttrs:{placeholder:"Enter here...",maxlength:300,value:this.note.text},trapFocus:!0,onConfirm:n=>this.editCurrent(n)})},editCurrent(n){n!==this.note.text&&this.editNoteAction({id:this.note.id,title:this.note.title,text:n,userId:this.note.userId,important:this.note.important,tags:this.note.tags})},openFull(){this.$buefy.dialog.alert({title:this.note.title,message:this.note.text,confirmText:"Cool!"})}}};var m=function(){var t=this,e=t._self._c;return e("div",{staticClass:"mt-2 mb-5"},[e("div",{staticClass:"box mr-2 ml-2",class:t.note.important?"green":"",staticStyle:{height:"250px",width:"300px",cursor:"pointer"}},[e("div",{on:{click:t.openFull}},[e("b-field",[e("b-tag",{attrs:{type:"is-primary","aria-close-label":"Close tag"}},[t._v(" "+t._s(t.note.tags)+" ")])],1),e("h3",{staticClass:"title is-4 mb-5"},[t._v(t._s(t.note.title))]),e("div",{staticClass:"mt-5",staticStyle:{"min-height":"90px"}},[t._v(" "+t._s(t.getBody())+" ")])],1),e("div",{staticStyle:{"margin-left":"160px"}},[e("b-button",{attrs:{type:"is-link","icon-right":"file-edit"},on:{click:function(s){return t.prompt()}}}),e("b-button",{staticClass:"ml-1",attrs:{type:"is-danger","icon-right":"delete"},on:{click:function(s){return t.deleteCurrent()}}})],1)])])},_=[],u=i(d,m,_,!1,null,"77b0c8a3",null,null);const h=u.exports;const f={name:"NoteList",components:{RegisterButton:r,LoginButton:l,NoteCard:h},computed:{...c(["allNotes","getNotImportant","getImportant"]),...p(["notes","token","search"])}};var g=function(){var t=this,e=t._self._c;return t.token!==null?e("div",{staticClass:"container mt-5"},[t.search!==""?e("div",{staticStyle:{display:"flex","flex-direction":"row","flex-wrap":"wrap",transform:"translateX(37px)"}},t._l(this.allNotes.filter(s=>s.tags===t.search||s.tags.startsWith(t.search)),function(s){return e("note-card",{attrs:{note:s}})}),1):t._e(),t.search===""?e("div",{staticStyle:{display:"flex","flex-direction":"row","flex-wrap":"wrap",transform:"translateX(37px)"}},t._l(this.getImportant,function(s){return e("note-card",{attrs:{note:s}})}),1):t._e(),this.getNotImportant.length!==0&&this.getImportant.length!==0?e("div",{staticClass:"text-divider"}):t._e(),t.search===""?e("div",{staticStyle:{display:"flex","flex-direction":"row","flex-wrap":"wrap",transform:"translateX(37px)"}},t._l(this.getNotImportant,function(s){return e("note-card",{attrs:{note:s}})}),1):t._e()]):e("div",{staticClass:"container",staticStyle:{"margin-top":"200px"}},[e("img",{staticClass:"mt-6",staticStyle:{margin:"0 auto",display:"block"},attrs:{src:"https://i.ibb.co/fS2mdmT/Dqv-Hnq-noteik-logo.png",alt:"Notik"}}),e("div",{staticStyle:{display:"flex","justify-content":"center","margin-top":"70px"}},[e("login-button"),e("register-button"),e("a",{staticClass:"button",staticStyle:{"margin-left":"10px"},attrs:{href:"https://github.com/gavrylenkoIvan",target:"_blank"}},[e("b-icon",{attrs:{icon:"github"}})],1)],1)])},v=[],x=i(f,g,v,!1,null,"029676e9",null,null);const b=x.exports;export{b as default};
