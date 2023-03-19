package teambitionapi

import "time"

var LocalTokenCache map[string]Token = make(map[string]Token)

type Token struct {
	tenantId   string
	token      string
	expireTime int64
}

type TokenCacheExecutor interface {
	SetToken(tenantId, token string, expireTime int64)
	GetToken(tenantId string) (token string, ok bool)
}

type DefaultTokenCacheExecutor struct {
}

func (d *DefaultTokenCacheExecutor) SetToken(tenantId, token string, expireTime int64) {
	LocalTokenCache[tenantId] = Token{
		tenantId:   tenantId,
		token:      token,
		expireTime: expireTime,
	}
}

func (d *DefaultTokenCacheExecutor) GetToken(tenantId string) (token string, ok bool) {
	t, ok := LocalTokenCache[tenantId]
	if !ok {
		return "", false
	}

	if t.expireTime < time.Now().Unix() {
		return "", false
	}
	return t.token, true
}
