**Frontend semplice con al limite un paio di librerie javascript**


Per creare una WPA è necessario un service worker che implementi il metodo fetch
=> per registrare un service worker è necessario un certificato SSL/TLS di cui il browser si fidi
==> nel caso in cui sia self-signed bisogna aggiungere quel certificato al browser dopo ogni aggiornamento
===> deve corrispondere all'URL servito dal web server => file hosts su windows per matchare IP con URL servito

Inoltre deve servire un icona quadrata di almeno 144x144 pixel, basta dargliela come impostazione nel manifest e non deve essere effettivamente cosi
Errori del manifest nella sezione Application dei Dev tools di chrome 

utilizzo il protocollo WebSocket per aggiungere la funzionalità di live reload: quando si interrompe la connessione ricarico la pagina