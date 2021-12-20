package main

type log4jFingerprint struct {
	file    string
	sha256  string
	version string
}

var log4jFingerprints = []log4jFingerprint{
	{file: "org/apache/logging/log4j/core/Filter.class", sha256: "28494eae67f5aad31597a7c883ea3a485e59983acbddf8c34931d038ae1880e5", version: "2.15.0"},
	{file: "org/apache/logging/log4j/core/LoggerContext.class", sha256: "b730113a53921e597a68dedd148977787ab1c6f045844868dbaa6e0d89e2e518", version: "2.14.1"},
	{file: "org/apache/logging/log4j/core/LoggerContext.class", sha256: "7875ed21f5b41de8f50df3d6ccf406c1b51432473f9c8e7fd5a6476a42aa13b8", version: "2.14.0"},
	{file: "org/apache/logging/log4j/core/appender/AbstractManager.class", sha256: "e1c1408b4564593d06016e84a17ec43031e10c36e1083d0538305f438d491517", version: "2.13.1"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "8c3ddc94600aa8696077fd67dc6cda00af740c96b047c2cfee0a333e04ef5b60", version: "2.12.1"},
	{file: "org/apache/logging/log4j/core/LoggerContext.class", sha256: "41ef81bc32b51ac9637f313f1140c1898409a23636531811bfc746315b731d16", version: "2.12.0"},
	{file: "org/apache/logging/log4j/core/LoggerContext.class", sha256: "c276ceae01fa13d9ee8f57fd5f4b2861b0ae36e3d16222a0504ee2ed9154c080", version: "2.11.2"},
	{file: "org/apache/logging/log4j/core/appender/AsyncAppender.class", sha256: "b5825fde082dcb6a05f15b3c6f514c9f73f7f8f9b95ec8da9d9a7947da918af4", version: "2.11.1"},
	{file: "org/apache/logging/log4j/core/AbstractLogEvent.class", sha256: "08e44a7a56082e3e65c8b6970fd3bb85cd067264dea6c7e0e6656d086c8967de", version: "2.11.0"},
	{file: "org/apache/logging/log4j/core/appender/ConsoleAppender.class", sha256: "20fbcf52f3fdd2058deedd12ca3c495f7c00caef5a863d8f85a61c9a4b2e8357", version: "2.10.0"},
	{file: "org/apache/logging/log4j/core/appender/AsyncAppender.class", sha256: "64b35cf90cd9d5e3a66643928bffbd7fc8f92e12d115d57a67a72acbcb214a22", version: "2.9.1"},
	{file: "org/apache/logging/log4j/core/appender/mom/JmsManager.class", sha256: "0260be3bb7a6c3bb5fa3792d2827e5bf240a90afb4dcccd72757ce19fa2dfdf0", version: "2.9.0"},
	{file: "org/apache/logging/log4j/core/LoggerContext.class", sha256: "54e97ca71b5250ad5c2b5f915c77281d542e667bfd7faed030f7f737f30bf170", version: "2.8.2"},
	{file: "org/apache/logging/log4j/core/appender/MemoryMappedFileManager.class", sha256: "de63748c77f86891f4909b8a11b3eebc5d4b8559ab42824b81316e0ae5d90c42", version: "2.8.1"},
	{file: "org/apache/logging/log4j/core/LoggerContext.class", sha256: "13a5f4e5de929b9f2506a9f10747729f39294dbca528bb865da216050948c6a4", version: "2.8"},
	{file: "org/apache/logging/log4j/ThreadContextAccess.class", sha256: "c8b3e6cb25b261a796b1fd5d6cd23dbfe1b1ef6d21bd9daf08e0cc77497ab92a", version: "2.7"},
	{file: "org/apache/logging/log4j/core/LoggerContext.class", sha256: "dcda8fb83531e52ab417106640554815d91e0571d2d8b9cd9cc29d3e071dda45", version: "2.6.2"},
	{file: "org/apache/logging/log4j/core/appender/OutputStreamManager.class", sha256: "387abe5220e2d26bf36c1037ebf343d0693ae304f7ad537f8c5f29352c53cd04", version: "2.6.1"},
	{file: "org/apache/logging/log4j/core/appender/OutputStreamManager.class", sha256: "f1c657573cf676bf14f63d102e71d481f32968c139e27359f447a4583eb6e29d", version: "2.6"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "dc6bbc06f212f040146711ad8b271a8d85b135f6504fd0f482086cf00e570da6", version: "2.5"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "fbf8b2f9b621e3cf9cef933403ea5b1b514d172cd7e32fc792d7e9a7c3a4eb2e", version: "2.4.1"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "9451fb5e697e44bb7f57846f6f10283bcbced5c5d00ddfdc102e90f66e3b229e", version: "2.4"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "f356fd80bc98ff8f41becae91f80c52149ecef1289db95b3f7c5cfb33facaa59", version: "2.3"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "9526d450bb5d1997295f5c3cdda79b007c49995f6ff0e412725205942bae36b0", version: "2.2"},
	{file: "org/apache/logging/log4j/core/AbstractLifeCycle.class", sha256: "774541fb268b292c69a72d0476e0738a1bcb6b51dc77d9d82602805b7c670183", version: "2.1"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "53785598c65f0db6ea42f5a6ec608f6628e605e3fe5df958909037f914427e9c", version: "2.0.2"},
	{file: "org/apache/logging/log4j/core/impl/Log4jLogEvent.class", sha256: "9a23e7642fd7b7dca23013d13cda4db91044f793101c2b9e030564da47cdab42", version: "2.0.1"},
	{file: "org/apache/logging/log4j/core/selector/JndiContextSelector.class", sha256: "d80fa3806065ccfcc0d0c678efe5c756ade1f8a5a72df46dcaed97e571fd7639", version: "2.0"},
	{file: "org/apache/logging/log4j/core/LogEventListener.class", sha256: "52f9c2741dc9fa16ca50aabd61d535624f5fefbbd5d1f88e963fd6a2dbc60717", version: "2.0-rc2"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "5ec127f5ab34c867e82b078431c7b54cdce46b1541c846ae4ecdda8ecccf0adf", version: "2.0-rc1"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "0ab344d7f3a244e76786932979eacafae1778cb53d772627c01464224b1007f4", version: "2.0-beta9"},
	{file: "org/apache/logging/log4j/core/appender/ConsoleAppender.class", sha256: "57174e9835ba74b954271375aaa3dacd57b6976c73ae7e34dc500e86e2525433", version: "2.0-beta8"},
	{file: "org/apache/logging/log4j/core/appender/ConsoleAppender.class", sha256: "2152e50ed5ea5e9a821853c81cf884173a8d51989de4507fde50b6eaace6f81d", version: "2.0-beta7"},
	{file: "org/apache/logging/log4j/core/LoggerContext.class", sha256: "bdc993287ad181389488fa255c48c7f5f2a7add441088363ad87b7c05dcf23a2", version: "2.0-beta6"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "2308742242c03027def9d1e6c614fcfe5f5b74673fc9030718654f9aeeb34ef6", version: "2.0-beta5"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "bf65cd0855b9cdc8b644ec956e34b10dad605c86de3fe7d871bc1d3f2be3c7f6", version: "2.0-beta4"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "4be5c0f7ca3e026af331e3a9f65d4b7af7acc56a52c8c261a30a81061696558f", version: "2.0-beta3"},
	{file: "org/apache/logging/log4j/core/Logger.class", sha256: "da6ce58a6c9fbcb7cd7caf4b2f5c539cf5f8e9bd4d55864d90ed4c4de8015a1d", version: "2.0-beta2"},
	{file: "org/apache/logging/log4j/core/LoggerContext.class", sha256: "24e764b716c678c3dfb905ec730aab7278464d1226e80dcfb62ac829b0b1839e", version: "2.0-alpha2"},
	{file: "org/apache/logging/log4j/core/LoggerContext.class", sha256: "2a6831860b3ab7b54165821ecfd62eb258ea382aa8ae9091281a7c85325a2f37", version: "2.0-alpha1"},
	{file: "org/apache/log4j/ConsoleAppender.class", sha256: "d73ce43dac92d151288da89ff6d765f31fc785ba540b0d18620b73d74aac56d9", version: "1.2.17"},
	{file: "org/apache/log4j/LogXF.class", sha256: "eb02dceb9eb42d04126c30828b1259dfdf9f316f070d46de1343ca832e150031", version: "1.2.16"},
	{file: "org/apache/log4j/AsyncAppender.class", sha256: "56d8abb7d76f00f06396fb3ed20aaaa654cd329c5bbb470e087c85cf8f8cd590", version: "1.2.15"},
	{file: "org/apache/log4j/Dispatcher.class", sha256: "f710cd2b3fdaa29fdbe683ed703a6ac41b4af66127c6b027996049fb65766390", version: "1.2.14"},
	{file: "org/apache/log4j/Dispatcher.class", sha256: "8c4a334ca28b5dae7915eb9c5a75ab740c3a198b7f68b8bcdf29220eb2f718b7", version: "1.2.13"},
	{file: "org/apache/log4j/ConsoleAppender.class", sha256: "612b80ce3c3e54d9c64d24395c09a94b109cf3424c2d3a906b3dfc240b0ac718", version: "1.2.12"},
	{file: "org/apache/log4j/ConsoleAppender.class", sha256: "000e92134b4aed946f96502207075c1de608305fa4952d5dcb2dd607ab064e06", version: "1.2.11"},
	{file: "org/apache/log4j/Logger.class", sha256: "96ea23a6b047028a003399a4454359ad61b38832e1c60d24d6d18b698d3891df", version: "1.2.9"},
	{file: "org/apache/log4j/ConsoleAppender.class", sha256: "24765dd3e6559df2d36b1bb2f3c9ab622bb84aab33fe76d5a0a3113b01daf25a", version: "1.2.8"},
	{file: "org/apache/log4j/LogManager.class", sha256: "f78ec34b23ac14035f0034fa3ed00841b2b66d9f8389b265a4aa07f79744106b", version: "1.2.7"},
	{file: "org/apache/log4j/MDC.class", sha256: "fafdf46644b35e408217890dc40a512df2a75f5aa0437e01cfeb84f9be75bb40", version: "1.2.6"},
	{file: "org/apache/log4j/net/SocketNode.class", sha256: "ed5d53deb29f737808521dd6284c2d7a873a59140e702295a80bd0f26988f53a", version: "1.2.5"},
	{file: "org/apache/log4j/Dispatcher.class", sha256: "ea5d8dac5f550597a914f07a43bcff51f934817412554b54d13a945504c91ef3", version: "1.2.4"},
	{file: "org/apache/log4j/ConsoleAppender.class", sha256: "1cfb24a5822a73d089baea464c2bc0bfe16da5b582bf323af8c813818529a3a6", version: "1.1.3"},
}