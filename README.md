# totp
simple TOTP implementation POC using following params:

X = 30, T0 = 0

TOTP = HOTP(K, T) More specifically, T = (Current Unix time - T0) / X, 

where the default floor function is used in the computation.

HOTP(K,C) = Truncate(HMAC-SHA-512(K,C))
