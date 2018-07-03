

var user={

    'ID': 0,
    'Login': '',
    'Password': '',
    'Username': '',
    'Email': '',
    'Status': false,
    'UserIcon': '',
};

var message={
    'User': user,
    'Group': {'GroupName': ''},
    'Content': '',
};
var group={
    'GroupName': '',
    'Messages':[message],
    'Members': [user],
    'ID': null,
};


var test = new Vue({
    el: '#app',
    data: {
        MessageIn:{
            User: user,
            Contact: user,
            Group: group,
            Message: message,
            Members: [user],
            RelationType: 0,
            MessageLimit: 0,
            Action: '',
        },

        User: user,
        ContactList: [user],
        GroupList: [group],
        Message: [message],
        Recipients: [user],
        Status: null,
        Action: '',
        ws: null, // Our websocket
        UsersFromServer: {},
        RecContents: {},
        RecContent: '',
        joined: false, // True if email and username have been filled in
        OnlineUsers: '',
        MyGroups: '',
    },


    created: function () {
        var self = this;
        var element = document.getElementById('chat-messages');
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function (e) {
            var msg = JSON.parse(e.data);

            if (msg.Action == "LoginUser") {
                if (typeof msg.User != "undefined") {
                    test.User = msg.User;
                }
                if(typeof msg.GroupList != "undefined" && msg.GroupList != null) {
                    test.GroupList = msg.GroupList;
                    for (var i = 0; i < msg.GroupList.length; i++) {
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
                        if(typeof msg.GroupList[i].Messages !=  "undefined") {
                            for (var j = 0; j <msg.GroupList[i].Messages.length; j++) {
                                if (typeof self.RecContents[msg.GroupList[i].GroupName] == "undefined") {
                                    self.RecContents[msg.GroupList[i].GroupName] = '';
                                }
                                self.RecContents[msg.GroupList[i].GroupName] +=
                                    '<div class="chip">' +
                                    msg.GroupList[i].Members[0].Username +
                                    '</div>' +
                                    '<div class="white-text">' +
                                    msg.GroupList[i].Messages[j].Content + '</div>' +
                                    '<br/>';
                            }
                        }
                    }
                }
            }else if (msg.Action == "SendMessageTo") {
                if(msg.Message.User.Username != test.User.Username) {
                    if (typeof self.RecContents[msg.Message.Group.GroupName] == "undefined") {
                        self.RecContents[msg.Message.Group.GroupName] = '';
                        var a = document.getElementById(test.User.Login + msg.Message.User.Login);
                        console.log(a);
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
                    }

                    self.RecContents[msg.Message.Group.GroupName] +=
                        '<div class="chip">' +
                        msg.Message.User.Username +
                        '</div>' +
                        '<div class="white-text">' +
                        msg.Message.Content + '</div>' +
                        '<br/>';

                    self.RecContent = self.RecContents[msg.Message.Group.GroupName];
                }
            }else if(msg.Action =="GetUsers"){
                //this.User = msg.User;
                self.OnlineUsers = '';
                if(typeof msg.ContactList != "undefined") {
                    for (var i = 0; i < msg.ContactList.length; i++) {
                        if (test.User.Login != msg.ContactList[i].Login) {
                            var gName = test.User.Login + msg.ContactList[i].Login;
                            //var rgName = msg.ContactList[i].Login+test.User.Login;
                            // var dex = true;
                            // if(typeof test.GroupList != "undefined" && test.GroupList != null) {
                            //     for(var g =0; g< test.GroupList.length;g++){
                            //         if(gName == test.GroupList[g] || rgName == test.GroupList[g]){
                            //             dex = false;
                            //         }
                            //     }
                            // }
                            test.UsersFromServer[gName] = msg.ContactList[i];
                            self.OnlineUsers +=
                                '<div class="input-field col s12">' +
                                '<button class="waves-effect waves-light btn col s12" onclick=createGroup(this) id = ' +
                                gName + '>' +
                                msg.ContactList[i].Username +
                                '</button></div>' +
                                '<br/>';

                        }
                    }
                }




            }

            element.scrollTop = element.scrollHeight;// Auto scroll to the bottom
        });
    },

    methods: {
        send: function () {
            if (this.MessageIn.Message.Content != '') {
                this.MessageIn.Message.User.Username = this.User.Username;
                this.MessageIn.Message.User.Login = this.User.Login;
                this.MessageIn.Message.User.ID = test.User.ID;
                this.MessageIn.Message.MessageSenderID = test.User.ID;
                this.MessageIn.User.Login = this.User.Login;
                this.MessageIn.Message.Group.MessageSenderID = test.User.ID;
                this.MessageIn.User.Username = this.User.Username;
                this.MessageIn.Message.Content = $('<p>').html(this.MessageIn.Message.Content).text();
                this.MessageIn.Message.Group.GroupName = this.MessageIn.Group.GroupName;
                this.MessageIn.Action = "SendMessageTo";
                if (typeof this.RecContents[this.MessageIn.Group.GroupName] == "undefined"){
                    this.RecContents[this.MessageIn.Group.GroupName] = '';
                }
                this.RecContents[this.MessageIn.Group.GroupName] +=
                    '<div class="chip">' +
                    test.User.Username +
                    '</div>' +
                    '<div class="white-text">' +
                    this.MessageIn.Message.Content +
                    '</div>' +
                    '<br/>';
                var self = this;
                self.RecContent = this.RecContents[this.MessageIn.Group.GroupName];
                this.ws.send(JSON.stringify(this.MessageIn));
                this.MessageIn.Message.Content = '';
            }
        },

        join: function () {
            if (!this.MessageIn.User.Login) {
                Materialize.toast('You must choose a username', 2000);
                return
            }

            this.MessageIn.User.Login = $('<p>').html(this.MessageIn.User.Login).text();
            this.MessageIn.User.Password = $('<p>').html(this.MessageIn.User.Password).text();
            this.MessageIn.User.Status = true;

            this.MessageIn.Action = "LoginUser";

            this.User.Username = this.MessageIn.User.Username;
            this.User.Login = this.MessageIn.User.Login;

            this.ws.send(JSON.stringify(this.MessageIn));
            this.joined = true;
        },
        signUp: function () {
            if (!this.MessageIn.User.Login) {
                Materialize.toast('You must choose a username', 2000);
                return
            }

            this.MessageIn.User.Username = $('<p>').html(this.MessageIn.User.Username).text();
            this.MessageIn.User.Login = $('<p>').html(this.MessageIn.User.Login).text();
            this.MessageIn.User.Password = $('<p>').html(this.MessageIn.User.Password).text();
            this.MessageIn.User.Email = $('<p>').html(this.MessageIn.User.Email).text();
            this.MessageIn.User.Status = true;
            this.MessageIn.User.UserIcon = "usericon";

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
    }
});

function changeUser(el) {
    test.MessageIn.Group.GroupName =el.id;
    test.RecContent = test.RecContents[el.id];
}

function createGroup(el) {
    var net = true;
    var rgName = test.UsersFromServer[el.id].Login+test.User.Login;
    if(typeof test.GroupList != "undefined" || test.GroupList !=null){
        for (var i = 0; i < test.GroupList.length; i++) {
            if (test.GroupList[i].GroupName == el.id || rgName == el.id) {
                net = false;
                break;
            }
        }
    }
    if (net){
        test.MessageIn.Action = "CreateGroup"
        test.MessageIn.Group.GroupTypeID = 1;
        test.MessageIn.Group.User = test.User;
        test.MessageIn.Group.GroupOwnerID = test.User.ID;
        test.MessageIn.Group.GroupName = el.id;
        test.MessageIn.Members[0] = test.User;
        test.MessageIn.Members[1] = test.UsersFromServer[el.id];
        test.ws.send(JSON.stringify(test.MessageIn))
        el.remove();
        var a = document.getElementById(el.id);
        if (a == null) {
            var element = document.getElementById('chat-messages');
            element.innerHTML += '<div class="input-field col s12">' +
                '<button class="waves-effect waves-light btn col s12" onclick=changeUser(this) id = ' +
                test.MessageIn.Group.GroupName + '>' +
                test.MessageIn.Members[1].Username +
                '</button></div>' +
                '<br/>';
        }
    }
}
