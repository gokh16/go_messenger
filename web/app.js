var humble = new Vue({
    el: '#app',
    data: {
        MessageIn:{
            User: {},
            Contact: {},
            Group: {},
            Message: {User: {}, Group: {}},
            Members: [],
            RelationType: 0,
            MessageLimit: 0,
            Action: '',
        },
        ProfileStr:{
            ProfileUser: {},
            Friend: false,
            NotMy: false,
            ProfileGroupName: '',
        },
        User: {},
        ContactList: [],
        GroupList: [],
        Message: [],
        Status: null,
        Action: '',
        ws: null,
        UsersFromServer: {},
        RecContents: {},
        RecContent: '',
        joined: false,
        profile: false,
        MyGroups: '<h4 class="white-text">Group list:</h4>',
        creatingGroup:'',
    },

    created: function () {
        var self = this;
        var element = document.getElementById('chat-messages');
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function (e) {
            var msg = JSON.parse(e.data);
            switch (msg.Action){
                case "LoginUser":
                    if(msg.Status == false){
                        alert("Не правильный пароль или логин!")
                        location.reload();
                    }
                    if (typeof msg.User != "undefined") {
                        this.User = msg.User;
                        humble.User = msg.User;
                    }
                    if (typeof msg.ContactList != "undefined"){
                        humble.ContactList = msg.ContactList;
                    }
                    if(typeof msg.GroupList != "undefined" && msg.GroupList != null) {
                        humble.GroupList = msg.GroupList;
                        for (var i = 0; i < msg.GroupList.length; i++) {
                            if(msg.GroupList[i].Members.length >2){
                                self.MyGroups +=
                                    '<div class="input-field col s12">' +
                                    '<button class="waves-effect waves-light btn col s12" onclick=changeUser(this) id = ' +
                                    msg.GroupList[i].GroupName + '>' +
                                    msg.GroupList[i].GroupName +
                                    '</button></div>' +
                                    '<br/>';
                            } else if(msg.GroupList[i].Members.length == 2) {
                                for (var c = 0; c < msg.GroupList[i].Members.length; c++) {
                                    if (msg.GroupList[i].Members[c].Login != self.User.Login) {
                                        self.MyGroups +=
                                            '<div class="input-field col s12">' +
                                            '<button class="waves-effect waves-light btn col s12" onclick=changeUser(this) id = ' +
                                            msg.GroupList[i].GroupName + '>' +
                                            msg.GroupList[i].Members[c].Username +
                                            '</button></div>' +
                                            '<br/>';
                                    }

                                }
                            }
                            if(typeof msg.GroupList[i].Messages !=  "undefined") {
                                for (var j = 0; j <msg.GroupList[i].Messages.length; j++) {
                                    var sen;
                                    if (typeof self.RecContents[msg.GroupList[i].GroupName] == "undefined") {
                                        self.RecContents[msg.GroupList[i].GroupName] = '';
                                    }
                                    for(var m=0;m<msg.GroupList[i].Members.length;m++){
                                        if(msg.GroupList[i].Members[m].ID == msg.GroupList[i].Messages[j].MessageSenderID){
                                            sen = msg.GroupList[i].Members[m];
                                        }
                                    }
                                    self.RecContents[msg.GroupList[i].GroupName] +=
                                        '<div class="chip">' +
                                        sen.Username +
                                        '</div>' +
                                        '<div class="white-text">' +
                                        msg.GroupList[i].Messages[j].Content + '</div>' +
                                        '<br/>';
                                }
                            }
                        }
                    }
                    break;
                case "SendMessageTo":
                    if(msg.Message.User.Username != humble.User.Username) {
                        if (typeof self.RecContents[msg.Message.Group.GroupName] == "undefined") {
                            self.RecContents[msg.Message.Group.GroupName] = '';
                            if(msg.Message.Group.GroupTypeID == 1) {
                                var a = document.getElementById(humble.User.Login + msg.Message.User.Login);
                                if (a != null) {
                                    a.remove();
                                }
                                self.MyGroups +=
                                    '<div class="input-field col s12">' +
                                    '<button class="waves-effect waves-light btn col s12" onclick=changeUser(this) id = ' +
                                    msg.Message.Group.GroupName + '>' +
                                    msg.Message.User.Username +
                                    '</button></div>' +
                                    '<br/>';
                            }else{
                                var elg = document.getElementById(msg.Message.Group.GroupName);
                                if(elg ==null) {
                                    self.MyGroups +=
                                        '<div class="input-field col s12">' +
                                        '<button class="waves-effect waves-light btn col s12" onclick=changeUser(this) id = ' +
                                        msg.Message.Group.GroupName + '>' +
                                        msg.Message.Group.GroupName +
                                        '</button></div>' +
                                        '<br/>';
                                }
                            }
                        }
                        if (Notification.permission === "granted") {
                            var notification = new Notification(msg.Message.User.Username+"\n"+msg.Message.Content);
                            setTimeout(notification.close.bind(notification), 1000);
                            notification = null;
                        }
                        self.RecContents[msg.Message.Group.GroupName] +=
                            '<div class="chip">' +
                            msg.Message.User.Username +
                            '</div>' +
                            '<div class="white-text">' +
                            msg.Message.Content + '</div>' +
                            '<br/>';
                        self.RecContent = self.RecContents[msg.Message.Group.GroupName];
                        this.scrollInBot();
                    }
                    break;
                case "GetUsers":
                    var modWin = document.getElementById('menuContent');
                    modWin.style.zIndex = '100';
                    modWin.style.display = 'block';
                    modWin.innerHTML ='<h4 class="white-text">User list:</h4>';
                    if (typeof msg.ContactList != "undefined") {
                        humble.ContactList = msg.ContactList;
                        for (var i = 0; i < msg.ContactList.length; i++) {
                            if (humble.User.Login != msg.ContactList[i].Login) {
                                var gName = humble.User.Login + msg.ContactList[i].Login;
                                humble.UsersFromServer[gName] = msg.ContactList[i];
                                modWin.innerHTML +=
                                    '<div class="input-field col s12">' +
                                    '<button class="waves-effect waves-light btn col s12" onclick=reftoshowProfile(this) id = ' +
                                    gName + '>' +
                                    msg.ContactList[i].Username +
                                    '</button></div>' +
                                    '<br/>';
                            }
                        }
                    }
                    break;
                case "GetContactList":
                    var modWin = document.getElementById('menuContent');
                    modWin.style.zIndex = '100';
                    modWin.style.display = 'block';
                    modWin.innerHTML ='<h4 class="white-text">Contact list:</h4>';
                    if (typeof msg.ContactList != "undefined") {
                        humble.ContactList = msg.ContactList;
                        for (var i = 0; i < msg.ContactList.length; i++) {
                            if (humble.User.Login != msg.ContactList[i].Login) {
                                var gName = humble.User.Login + msg.ContactList[i].Login;
                                humble.UsersFromServer[gName] = msg.ContactList[i];
                                modWin.innerHTML +=
                                    '<div class="input-field col s12">' +
                                    '<button class="waves-effect waves-light btn col s12" onclick=reftoshowProfile(this) id = ' +
                                    gName + '>' +
                                    msg.ContactList[i].Username +
                                    '</button></div>' +
                                    '<br/>';
                            }
                        }
                    }
                    break;
            }

        });
    },

    methods: {
        send: function () {
            var element = document.getElementById('chat-messages');
            if (this.MessageIn.Message.Content != '' && typeof this.MessageIn.Group.GroupName !="undefined") {
                this.MessageIn.Message.User.Username = humble.User.Username;
                this.MessageIn.Message.User.Login = humble.User.Login;
                this.MessageIn.Message.User.ID = humble.User.ID;
                this.MessageIn.Message.MessageSenderID = humble.User.ID;
                this.MessageIn.User.Login = this.User.Login;
                this.MessageIn.Message.Group.MessageSenderID = humble.User.ID;
                this.MessageIn.User.Username = this.User.Username;
                this.MessageIn.Message.Content = $('<p>').html(this.MessageIn.Message.Content).text();
                this.MessageIn.Message.Group.GroupName = this.MessageIn.Group.GroupName;
                if(typeof humble.GroupList != "undefined")
                    for(var i=0;i<humble.GroupList.length;i++)
                    {
                        if(humble.GroupList[i].GroupName == this.MessageIn.Group.GroupName)
                        {
                            this.MessageIn.Message.Group = humble.GroupList[i];
                            this.MessageIn.Message.MessageRecipientID=humble.GroupList[i].ID;
                        }
                    }
                this.MessageIn.Action = "SendMessageTo";
                if (typeof this.RecContents[this.MessageIn.Group.GroupName] == "undefined"){
                    this.RecContents[this.MessageIn.Group.GroupName] = '';
                }
                this.RecContents[this.MessageIn.Group.GroupName] +=
                    '<div class="chip">' +
                    humble.User.Username +
                    '</div>' +
                    '<div class="white-text">' +
                    this.MessageIn.Message.Content +
                    '</div>' +
                    '<br/>';
                var self = this;
                self.RecContent = this.RecContents[this.MessageIn.Group.GroupName];
                this.ws.send(JSON.stringify(this.MessageIn));
                this.MessageIn.Message.Content = '';
                this.$nextTick(function () {
                    element.scrollTop = element.scrollHeight;
                })

            }
        },

        join: function () {
            if (Notification.permission !== 'denied') {
                Notification.requestPermission();
            }
            if (!this.MessageIn.User.Login) {
                Materialize.toast('You must choose a login', 2000);
                return;
            }
            if (!this.MessageIn.User.Password) {
                Materialize.toast('You must choose a password', 2000);
                return;
            }

            this.MessageIn.User.Login = $('<p>').html(this.MessageIn.User.Login).text();
            this.MessageIn.User.Password = $('<p>').html(this.MessageIn.User.Password).text();
            this.MessageIn.User.Status = true;
            this.MessageIn.Action = "LoginUser";
            this.User.Login = this.MessageIn.User.Login;
            this.ws.send(JSON.stringify(this.MessageIn));
            document.title = this.User.Login;
            this.joined = true;
        },

        signUp: function () {
            if (!this.MessageIn.User.Login) {
                Materialize.toast('You must choose a login', 2000);
                return;
            }
            if (!this.MessageIn.User.Username) {
                Materialize.toast('You must choose a username', 2000);
                return;
            }
            if (!this.MessageIn.User.Password) {
                Materialize.toast('You must choose a password', 2000);
                return;
            }
            if (!this.MessageIn.User.Email) {
                Materialize.toast('You must choose a email', 2000);
                return;
            }

            this.MessageIn.User.Username = $('<p>').html(this.MessageIn.User.Username).text();
            this.MessageIn.User.Login = $('<p>').html(this.MessageIn.User.Login).text();
            this.MessageIn.User.Password = $('<p>').html(this.MessageIn.User.Password).text();
            this.MessageIn.User.Email = $('<p>').html(this.MessageIn.User.Email).text();
            this.MessageIn.User.UserIcon = $('<p>').html(this.MessageIn.User.UserIcon).text();
            if(this.MessageIn.User.UserIcon == ''){
                this.MessageIn.User.UserIcon = "http://www.it-academy.kg/img/lang_img/hand/750e3525a8c1f3f242c2a7a081ef464d.png";
            }
            this.MessageIn.User.Status = true;
            this.MessageIn.Action = "CreateUser";
            this.User.Username = this.MessageIn.User.Username;
            this.User.Login = this.MessageIn.User.Login;
            this.ws.send(JSON.stringify(this.MessageIn));
            location.href="index.html"
        },

        showUsers: function (){
            this.MessageIn.User.Login = this.User.Login;
            this.MessageIn.User.Username = this.User.Username;
            this.MessageIn.Action = "GetUsers";
            this.ws.send(JSON.stringify(this.MessageIn))
        },

        showGroupListMob:function(){
            var modWin = document.getElementById('menuContent');
            modWin.style.display = 'none';
            var el =document.getElementById("chat-messages");
            el.style += 'z-index: 120';

            el.innerHTML = '';
            el.innerHTML+=this.MyGroups;
        },

        createPublicGroup: function(){
            var modWin = document.getElementById('menuContent');
            modWin.style.display = 'block';
            modWin.style.zIndex = '100';
            if (typeof this.ContactList != "undefined") {
                modWin.innerHTML = ' <div class="input-field white-text "><input type="text" id="creatingGroup">\n' +
                    '<label for="creatingGroup">Название группы</label></div>';
                for (var i = 0; i < this.ContactList.length; i++) {
                    if (this.User.Login != this.ContactList[i].Login) {
                        modWin.innerHTML +=
                            '<div class="form-check">' +
                            '<input type="checkbox" class="form-check-input "  id = ' +
                            this.ContactList[i].Login + '><label class="form-check-label" for=' +
                            this.ContactList[i].Login + '>'+
                            this.ContactList[i].Username +'</label>'+
                            '</div>' +
                            '<br/>';
                    }
                }
            }
            modWin.innerHTML+= '<div class="input-field col s12">' +
                '<button class="waves-effect waves-light btn col s12" onclick=createPubGroup()>' +
                'Создать группу' +
                '</button></div>' +
                '<br/>';


        },

        search: function(){
            this.MessageIn.Action = "GetUsers";
            this.ws.send(JSON.stringify(this.MessageIn))
        },

        burger: function () {
            $('.menu').toggleClass('menu_opened');
            $(document).click(function(event) {
                if ($(event.target).closest(".burger_trigger").length ) return;
                $('.menu').removeClass('menu_opened');
                event.stopPropagation();
            });
        },

        openProfile: function(){
            this.ProfileStr.ProfileUser = this.User;
            this.profile = true;
        },

        showContacts: function(){
            this.MessageIn.User.Login = this.User.Login;
            this.MessageIn.User.Username = this.User.Username;
            this.MessageIn.Action = "GetContactList";
            this.ws.send(JSON.stringify(this.MessageIn))
        },

        changeUserFromProfile:function(){
            this.backtochat();
            var el = this.ProfileStr.ProfileGroupName;
            var net = true;
            var rgName = this.UsersFromServer[el].Login+this.User.Login;
            if(typeof this.GroupList != "undefined" || this.GroupList !=null){
                for (var i = 0; i < this.GroupList.length; i++) {
                    if (this.GroupList[i].GroupName == el || rgName == this.GroupList[i].GroupName) {
                        if(rgName == this.GroupList[i].GroupName) el = rgName;
                        net = false;
                        break;
                    }
                }
            }
            if (net){
                this.MessageIn.Action = "CreateGroup"
                this.MessageIn.Group.GroupTypeID = 1;
                this.MessageIn.Group.User = this.User;
                this.MessageIn.Group.GroupOwnerID = this.User.ID;
                this.MessageIn.Group.GroupName = el;
                this.MessageIn.Members[0] = this.User;
                this.MessageIn.Members[1] = this.UsersFromServer[el];
                this.ws.send(JSON.stringify(this.MessageIn))
                this.$nextTick(function () {
                    var element = document.getElementById('groupList');
                    /*element.innerHTML += '<div class="input-field col s12">' +
                        '<button class="waves-effect waves-light btn col s12" onclick=changeUser(this) id = ' +
                        this.ProfileStr.ProfileGroupName + '>' +
                        this.UsersFromServer[this.ProfileStr.ProfileGroupName].Username +
                        '</button></div>' +
                        '<br/>';*/

                    this.MyGroups += '<div class="input-field col s12">' +
                        '<button class="waves-effect waves-light btn col s12" onclick=changeUser(this) id = ' +
                        this.ProfileStr.ProfileGroupName + '>' +
                        this.UsersFromServer[this.ProfileStr.ProfileGroupName].Username +
                        '</button></div>' +
                        '<br/>';
                })
            }else{
                this.$nextTick(function () {
                    console.log(el);
                    this.MessageIn.Group.GroupName =el ;
                    this.RecContent = this.RecContents[el];
                    this.scrollInBot();
                })
            }
        },

        AddContact:function(){
            this.MessageIn.Action = "AddContact";
            this.MessageIn.RelationType=1;
            this.MessageIn.Contact = this.ProfileStr.ProfileUser;
            humble.ws.send(JSON.stringify(humble.MessageIn));
        },

        backtochat: function(){
            this.profile = false;
        },

        showProfile: function(a){
            if(this.ContactList != "undefined"){
                for(var i=0;i<this.ContactList.length;i++){
                    if(this.ContactList[i].Login == this.UsersFromServer[a].Login){
                        this.Friend = true;
                    }
                }
            }
            this.profile =true;
            this.ProfileStr.ProfileUser = this.UsersFromServer[a];
            this.ProfileStr.NotMy=true;
            this.ProfileStr.ProfileGroupName = a;
        },

        scrollInBot: function(){
            var element = document.getElementById('chat-messages');
            this.$nextTick(function () {
                element.scrollTop = element.scrollHeight;

            })
        },

        exit: function () {
            this.joined = false;
            location.reload();
        }
    }
});

function changeUser(el) {
    humble.MessageIn.Group.GroupName =el.id;
    humble.RecContent = humble.RecContents[el.id];
    humble.scrollInBot();
}

function addGroupMember(el){
    humble.MessageIn.Action = "AddGroupMember"
    humble.ws.send(JSON.stringify(humble.MessageIn))
}
function createPubGroup() {
    var el1 = document.getElementById('creatingGroup');
    if (!el1.value) {
        Materialize.toast('You must write a group name', 2000);
        return;
    }
    var chbox=[];
    var cout =0;
    if (typeof humble.ContactList != "undefined") {
        for (var i = 0; i < humble.ContactList.length; i++) {
            chbox[cout]=document.getElementById(humble.ContactList[i].Login);
            if(chbox[cout] != null){
                cout++;
            }
        }
    }
    cout = 0;
    var accept = 0;
    document.getElementById('menuContent').innerHTML = '';
    for(var i =0;i<chbox.length; i++){
        if(chbox[i].checked){
            accept++;
            var user1 = {};
            user1.Login = chbox[i].id;
            console.log("sa",user1.Login);
            humble.MessageIn.Members[cout] = user1;
            //humble.MessageIn.Members[cout].Login = chbox[i].id;
            cout++;
        }
    }
    if(accept>0 && el1.value != null) {
        humble.MessageIn.Members[humble.MessageIn.Members.length] = humble.User;
        humble.MessageIn.Group.Members = humble.MessageIn.Members;

        humble.MessageIn.Action = "CreateGroup"
        humble.MessageIn.Group.GroupTypeID = 2;
        humble.MessageIn.Group.User = humble.User;
        humble.MessageIn.Group.GroupOwnerID = humble.User.ID;
        humble.MessageIn.Group.GroupName = el1.value;
        humble.ws.send(JSON.stringify(humble.MessageIn));
        var modWin = document.getElementById('chat-messages');
        modWin.innerHTML = '';
        var element = document.getElementById('groupList');
        element.innerHTML += '<div class="input-field col s12">' +
            '<button class="waves-effect waves-light btn col s12" onclick=changeUser(this) id = ' +
            humble.MessageIn.Group.GroupName + '>' +
            humble.MessageIn.Group.GroupName +
            '</button></div>' +
            '<br/>';
    }

}
function reftoshowProfile(el) {
    humble.showProfile(el.id);
}