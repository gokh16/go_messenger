

var user={
    'Login': '',
    'Password': '',
    'Username': '',
    'Email': '',
    'Status': false,
    'UserIcon': '',
};

var message={
    'Content': '',
};
var group={
    'GroupName': '',
    'Messages':[message],
    'Members': [user],
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
        Status: null,
        Action: '',
        ws: null, // Our websocket

        RecContents: {},
        RecContent: '',
        joined: false, // True if email and username have been filled in
        OnlineUsers: '',
        OurUsername: '',
    },


    created: function () {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function (e) {
            var msg = JSON.parse(e.data);
            if (msg.Action == "LoginUser") {
                for (var i = 0; i < msg.GroupList.length; i++) {
                    for (var c = 0; c < msg.GroupList[i].Members.length; c++) {
                        self.OnlineUsers +='<div class="input-field col s12"><button class="waves-effect waves-light btn col s12" onclick="changeUser(this)" id = "'+msg.GroupList[i].Members[c].Login+'">' + msg.GroupList[i].Members[c].Login + '</button></div>' + '<br/>';


                    }
                }
            } else if (msg.Action == "SendMessageTo") {
                RecContents[msg.User.Username] = msg.Message.Content;
                self.RecContent +=
                    '<div class="chip">' + msg.User.Username + ' from: ' + msg.GroupList[0].GroupName + '</div>' +
                    '<div class="white-text">' + msg.Message.Content + '</div>' +
                    '<br/>';
            }
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight;// Auto scroll to the bottom
        });
    },

    methods: {
        send: function () {
            if (this.MessageIn.Message.Content != '') {
                this.MessageIn.User.Login = this.OurUsername;
                this.MessageIn.User.Username = this.OurUsername;
                this.MessageIn.Message.Content = $('<p>').html(this.MessageIn.Message.Content).text();
                this.MessageIn.Group.GroupName= $('<p>').html(this.MessageIn.Group.GroupName).text();
                this.MessageIn.Action = "SendMessageTo";
                this.ws.send(JSON.stringify(this.MessageIn));
                this.MessageIn.Message.Content = ''; // Reset newMsg
            }
        },

        join: function () {
            if (!this.MessageIn.User.Login) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.MessageIn.User.Username = $('<p>').html(this.MessageIn.User.Login).text();
            this.MessageIn.User.Login = $('<p>').html(this.MessageIn.User.Login).text();
            this.MessageIn.User.Password = $('<p>').html(this.MessageIn.User.Password).text();
            this.MessageIn.User.Email = "email";
            this.MessageIn.User.Status = true;
            this.MessageIn.User.UserIcon = "usericon";
            this.MessageIn.Action = "LoginUser";
            this.OurUsername = $('<p>').html(this.MessageIn.User.Login).text();
            this.ws.send(JSON.stringify(this.MessageIn));
            this.joined = true;
        },
        signUp: function () {
            if (!this.MessageIn.User.Login) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.MessageIn.User.Username = $('<p>').html(this.MessageIn.User.Login).text();
            this.MessageIn.User.Login = $('<p>').html(this.MessageIn.User.Login).text();
            this.MessageIn.User.Password = $('<p>').html(this.MessageIn.User.Password).text();
            this.MessageIn.User.Email = "email";
            this.MessageIn.User.Status = true;
            this.MessageIn.User.UserIcon = "usericon";
            this.MessageIn.Action = "CreateUser";
            this.ws.send(JSON.stringify(this.MessageIn));
        },
        changeUser: function () {
            var self = this;
            self.RecContent +=
                '<div class="chip">' + $(this).attr('id') + '</div>' +
                '<br/>';
        },
    }
});

function changeUser(el) {

    test.RecContent = test.RecContents[el.id];
}

