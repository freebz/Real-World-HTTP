// 예제 5-8 파이프라이닝이 가능한지 판정하는 파이어폭스의 소스 코드

// the list of servers known to do bad things with pipelined requests
static const char *bad_servers[] = {
    "Microsoft-IIS/4.",
    "Microsoft-IIS/5.",
    "Netscape-Enterprise/3.",
    nsnull
};

for (const char **server = bad_servers; *server; ++server) {
    if (PL_strcasestr(val, *server) != nsnull) {
	LOG(("looks like this server does not support pipelining"));
	return PR_FALSE;
    }
}
