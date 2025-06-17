# Traefik Plugin Basic Auth Decoder

Ce plugin pour Traefik décode le header d'autorisation Basic Auth et ajoute le nom d'utilisateur dans un nouveau header `X-Traefik-Loggable-Username`.

## Installation

1. Compilez le plugin :
```bash
go build -o basic-auth-decoder.so -buildmode=plugin
```

2. Configurez Traefik pour utiliser le plugin en ajoutant la configuration suivante dans votre fichier de configuration Traefik :

```yaml
experimental:
  plugins:
    basicAuthDecoder:
      moduleName: github.com/gmarsay/traefik-plugin-basic-auth-decoder
      version: v0.1.0
```

## Utilisation

Ajoutez le middleware dans votre configuration de route :

```yaml
http:
  middlewares:
    basic-auth-decoder:
      plugin:
        basicAuthDecoder:
```

## Fonctionnement

Le plugin :
1. Vérifie la présence du header `Authorization`
2. Si présent et au format Basic Auth, décode le token
3. Extrait le nom d'utilisateur
4. Ajoute le nom d'utilisateur dans le header `X-Traefik-Loggable-Username`

## Licence

MIT