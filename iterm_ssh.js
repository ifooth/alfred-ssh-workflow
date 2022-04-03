#!/usr/bin/env osascript -l JavaScript
function run(argv) {
    console.log("argv: " + argv)
    console.log('argv length: ' + argv.length)

    if (argv.length == 0) {
        console.log("argv is empty, done");
        return
    }
    if (argv.length == 1 && argv[0].includes("\n")) {
        console.log("split by line");
        var inputs = argv[0].split("\n");
    } else {
        console.log("use shell argv");
        var inputs = argv;
    }
    _run(inputs)
}

function _run(inputs) {
    app = Application("iTerm");
    app.activate()

    if (app.windows.length == 0) {
        var window = app.createWindowWithDefaultProfile();
        var session = window.currentSession();
    } else {
        var window = app.currentWindow();
        var tab = window.createTabWithDefaultProfile();
        var session = tab.currentSession();
    }

    for (var i = 0; i < inputs.length; i++) {
        input = inputs[i]
        if (input.slice(0, 9) == 'password:') {
            handler_password(session, input.slice(9))
        }
        else if (input.slice(0, 9) == 'waitdone:') {
            handler_waitdone(session, input.slice(9))
        } else {
            handler_input(session, input)
        }

    }
    console.log("login done")

}

function handler_password(session, input) {
    var n = 1
    while (n < 20) {
        delay(1)

        console.log('try to login [' + n + ']')
        var contents = session.text().trim().split("\n");
        var lastLine = contents[contents.length - 1]

        if (lastLine.includes("assword:")) {
            session.write({ text: input });
            console.log("rows: " + session.rows())
            console.log("input password success");
            break;
        }
        if (lastLine.includes("#") || lastLine.includes("$")) {
            console.log("already login, ignore input password");
            break;
        }
        n = n + 1
    }
}

function handler_waitdone(session, input) {
    var n = 1
    while (n < 10) {
        delay(1)

        console.log('try to inspect input shell [' + n + ']')
        var contents = session.text().trim().split("\n");
        var lastLine = contents[contents.length - 1]

        if (lastLine.includes("#") || lastLine.includes("$") || lastLine.includes("~")) {
            session.write({ text: input });
            console.log("success get shell, input shell done");
            break;
        }
        n = n + 1
    }
}

function handler_input(session, input) {
    session.write({ text: input });
}