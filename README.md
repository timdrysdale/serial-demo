Working after using serial monitor (and not resetting the connection)
```
[timestamp] [threadID] facility level [function call] <message>
--------------------------------------------------------------------------------
[ 0.010873] [002b729b] libusb: debug [libusb_get_device_list] 
[ 0.010936] [002b729b] libusb: debug [libusb_get_device_descriptor] 
[ 0.010973] [002b729b] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.011002] [002b729d] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.011032] [002b729b] libusb: debug [parse_configuration] skipping descriptor 0xb
[ 0.011086] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011101] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 25
[ 0.011068] [002b729d] libusb: debug [handle_events] poll fds modified, reallocating
[ 0.011129] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011155] [002b729d] libusb: debug [handle_events] poll() 2 fds with timeout in 100ms
[ 0.011163] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011195] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011207] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011219] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011230] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011240] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011251] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011263] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011274] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011285] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011297] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011309] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011320] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011331] [002b729b] libusb: debug [parse_endpoint] skipping descriptor b
[ 0.011344] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011354] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 25
[ 0.011366] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011377] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 25
[ 0.011388] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011399] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 25
[ 0.011410] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011422] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 25
[ 0.011445] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011544] [002b729b] libusb: debug [libusb_get_device_descriptor] 
[ 0.011562] [002b729b] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.011576] [002b729b] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.011599] [002b729b] libusb: debug [libusb_get_device_descriptor] 
[ 0.011615] [002b729b] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.011635] [002b729b] libusb: debug [libusb_get_device_descriptor] 
[ 0.011649] [002b729b] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.011671] [002b729b] libusb: debug [libusb_open] open 1.13
[ 0.011721] [002b729b] libusb: debug [usbi_add_pollfd] add fd 9 events 4
[ 0.011746] [002b729b] libusb: debug [libusb_get_device_descriptor] 
[ 0.011762] [002b729b] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.011800] [002b729b] libusb: debug [libusb_get_device_descriptor] 
[ 0.011804] [002b729d] libusb: debug [handle_events] poll() returned 1
[ 0.011851] [002b729d] libusb: debug [handle_events] caught a fish on the event pipe
[ 0.011871] [002b729d] libusb: debug [handle_events] someone updated the poll fds
[ 0.011816] [002b729b] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.011902] [002b729d] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.011916] [002b729d] libusb: debug [handle_events] poll fds modified, reallocating
[ 0.011915] [002b729b] libusb: debug [libusb_get_device_descriptor] 
[ 0.011944] [002b729d] libusb: debug [handle_events] poll() 3 fds with timeout in 100ms
[ 0.011954] [002b729b] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.011985] [002b729b] libusb: debug [libusb_get_configuration] 
[ 0.012035] [002b729b] libusb: debug [libusb_get_configuration] active config 1
[ 0.012053] [002b729b] libusb: debug [libusb_detach_kernel_driver] interface 0
[ 0.012539] [002b729b] libusb: debug [libusb_get_configuration] 
[ 0.012591] [002b729b] libusb: debug [libusb_get_configuration] active config 1
[ 0.012603] [002b729b] libusb: debug [libusb_claim_interface] interface 0
[ 0.012642] [002b729b] libusb: debug [libusb_alloc_transfer] transfer 0x225f190
[ 0.012671] [002b729b] libusb: debug [libusb_submit_transfer] transfer 0x225f190
[ 0.012682] [002b729b] libusb: debug [submit_bulk_transfer] need 1 urbs for new transfer with length 28
[ 0.012823] [002b729d] libusb: debug [handle_events] poll() returned 1
[ 0.012865] [002b729d] libusb: debug [reap_for_handle] urb type=3 status=0 transferred=28
[ 0.012880] [002b729d] libusb: debug [handle_bulk_completion] handling completion status 0 of bulk urb 1/1
[ 0.012894] [002b729d] libusb: debug [handle_bulk_completion] last URB in transfer --> complete!
[ 0.012907] [002b729d] libusb: debug [usbi_handle_transfer_completion] transfer 0x225f190 has callback 0x4ad830
[ 0.013005] [002b729d] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.013024] [002b729d] libusb: debug [handle_events] poll() 3 fds with timeout in 100ms
[ 0.013013] [002b729b] libusb: debug [libusb_free_transfer] transfer 0x225f190
28 bytes successfully sent to the endpoint
INFO[0000] vid=0403,pid=6001,bus=1,addr=13,config=1,if=0,alt=0.InEndpoint(1) MaxPacketSize: 64 
[ 0.013325] [002b729b] libusb: debug [libusb_alloc_transfer] transfer 0x225ff40
[ 0.013341] [002b729b] libusb: debug [libusb_submit_transfer] transfer 0x225ff40
[ 0.013353] [002b729b] libusb: debug [submit_bulk_transfer] need 1 urbs for new transfer with length 192
[ 0.013432] [002b729d] libusb: debug [handle_events] poll() returned 1
[ 0.013452] [002b729d] libusb: debug [reap_for_handle] urb type=3 status=0 transferred=33
[ 0.013464] [002b729d] libusb: debug [handle_bulk_completion] handling completion status 0 of bulk urb 1/1
[ 0.013476] [002b729d] libusb: debug [handle_bulk_completion] last URB in transfer --> complete!
[ 0.013488] [002b729d] libusb: debug [usbi_handle_transfer_completion] transfer 0x225ff40 has callback 0x4ad830
[ 0.013517] [002b729d] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.013528] [002b729d] libusb: debug [handle_events] poll() 3 fds with timeout in 100ms
[ 0.013534] [002b729b] libusb: debug [libusb_free_transfer] transfer 0x225ff40
    [0000] read 33: {"report":"port","is":"open"}
[ 0.113743] [002b729d] libusb: debug [handle_events] poll() returned 0
[ 0.113796] [002b729d] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.113809] [002b729d] libusb: debug [handle_events] poll() 3 fds with timeout in 100ms
[ 0.114054] [002b729f] libusb: debug [libusb_alloc_transfer] transfer 0x7f7200000bd0
[ 0.114128] [002b729f] libusb: debug [libusb_submit_transfer] transfer 0x7f7200000bd0
[ 0.114145] [002b729f] libusb: debug [submit_bulk_transfer] need 1 urbs for new transfer with length 192
[ 0.114240] [002b729d] libusb: debug [handle_events] poll() returned 1
[ 0.114265] [002b729d] libusb: debug [reap_for_handle] urb type=3 status=0 transferred=33
[ 0.114279] [002b729d] libusb: debug [handle_bulk_completion] handling completion status 0 of bulk urb 1/1
[ 0.114293] [002b729d] libusb: debug [handle_bulk_completion] last URB in transfer --> complete!
[ 0.114308] [002b729d] libusb: debug [usbi_handle_transfer_completion] transfer 0x7f7200000bd0 has callback 0x4ad830
[ 0.114355] [002b729d] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.114370] [002b729d] libusb: debug [handle_events] poll() 3 fds with timeout in 100ms
[ 0.114380] [002b729f] libusb: debug [libusb_free_transfer] transfer 0x7f7200000bd0
    [0000] read 33: `{"report":"port","is":"open"}
```

Not working after replug without using serial monitor
```
[timestamp] [threadID] facility level [function call] <message>
--------------------------------------------------------------------------------
[ 0.008734] [002b73c9] libusb: debug [libusb_get_device_list] 
[ 0.008804] [002b73c9] libusb: debug [libusb_get_device_descriptor] 
[ 0.008842] [002b73c9] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.008900] [002b73c9] libusb: debug [parse_configuration] skipping descriptor 0xb
[ 0.008892] [002b73cb] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.008957] [002b73cb] libusb: debug [handle_events] poll fds modified, reallocating
[ 0.008914] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.008999] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 25
[ 0.008985] [002b73cb] libusb: debug [handle_events] poll() 2 fds with timeout in 100ms
[ 0.009024] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009037] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009048] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009059] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009070] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009080] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009091] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009101] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009113] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009124] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009135] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009146] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009157] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009168] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009178] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor b
[ 0.009191] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009201] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 25
[ 0.009212] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009222] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 25
[ 0.009234] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009244] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 25
[ 0.009255] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009265] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 25
[ 0.009276] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009375] [002b73c9] libusb: debug [libusb_get_device_descriptor] 
[ 0.009390] [002b73c9] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.009403] [002b73c9] libusb: debug [parse_endpoint] skipping descriptor 30
[ 0.009425] [002b73c9] libusb: debug [libusb_get_device_descriptor] 
[ 0.009439] [002b73c9] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.009457] [002b73c9] libusb: debug [libusb_get_device_descriptor] 
[ 0.009469] [002b73c9] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.009487] [002b73c9] libusb: debug [libusb_open] open 1.14
[ 0.009531] [002b73c9] libusb: debug [usbi_add_pollfd] add fd 9 events 4
[ 0.009559] [002b73c9] libusb: debug [libusb_get_device_descriptor] 
[ 0.009573] [002b73c9] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.009609] [002b73c9] libusb: debug [libusb_get_device_descriptor] 
[ 0.009623] [002b73c9] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.009620] [002b73cb] libusb: debug [handle_events] poll() returned 1
[ 0.009638] [002b73c9] libusb: debug [libusb_get_device_descriptor] 
[ 0.009691] [002b73c9] libusb: debug [libusb_get_config_descriptor] index 0
[ 0.009669] [002b73cb] libusb: debug [handle_events] caught a fish on the event pipe
[ 0.009729] [002b73c9] libusb: debug [libusb_get_configuration] 
[ 0.009731] [002b73cb] libusb: debug [handle_events] someone updated the poll fds
[ 0.009781] [002b73c9] libusb: debug [libusb_get_configuration] active config 1
[ 0.009780] [002b73cb] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.009804] [002b73c9] libusb: debug [libusb_detach_kernel_driver] interface 0
[ 0.009822] [002b73cb] libusb: debug [handle_events] poll fds modified, reallocating
[ 0.009867] [002b73cb] libusb: debug [handle_events] poll() 3 fds with timeout in 100ms
[ 0.010305] [002b73c9] libusb: debug [libusb_get_configuration] 
[ 0.010350] [002b73c9] libusb: debug [libusb_get_configuration] active config 1
[ 0.010366] [002b73c9] libusb: debug [libusb_claim_interface] interface 0
[ 0.010402] [002b73c9] libusb: debug [libusb_alloc_transfer] transfer 0x18e6190
[ 0.010430] [002b73c9] libusb: debug [libusb_submit_transfer] transfer 0x18e6190
[ 0.010442] [002b73c9] libusb: debug [submit_bulk_transfer] need 1 urbs for new transfer with length 28
[ 0.010537] [002b73cb] libusb: debug [handle_events] poll() returned 1
[ 0.010559] [002b73cb] libusb: debug [reap_for_handle] urb type=3 status=0 transferred=28
[ 0.010573] [002b73cb] libusb: debug [handle_bulk_completion] handling completion status 0 of bulk urb 1/1
[ 0.010586] [002b73cb] libusb: debug [handle_bulk_completion] last URB in transfer --> complete!
[ 0.010599] [002b73cb] libusb: debug [usbi_handle_transfer_completion] transfer 0x18e6190 has callback 0x4ad830
[ 0.010697] [002b73cb] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.010709] [002b73cb] libusb: debug [handle_events] poll() 3 fds with timeout in 100ms
[ 0.010774] [002b73c9] libusb: debug [libusb_free_transfer] transfer 0x18e6190
28 bytes successfully sent to the endpoint
INFO[0000] vid=0403,pid=6001,bus=1,addr=14,config=1,if=0,alt=0.InEndpoint(1) MaxPacketSize: 64 
[ 0.011088] [002b73c9] libusb: debug [libusb_alloc_transfer] transfer 0x18e6f40
[ 0.011104] [002b73c9] libusb: debug [libusb_submit_transfer] transfer 0x18e6f40
[ 0.011115] [002b73c9] libusb: debug [submit_bulk_transfer] need 1 urbs for new transfer with length 192
[ 0.011178] [002b73cb] libusb: debug [handle_events] poll() returned 1
[ 0.011212] [002b73cb] libusb: debug [reap_for_handle] urb type=3 status=0 transferred=2
[ 0.011229] [002b73cb] libusb: debug [handle_bulk_completion] handling completion status 0 of bulk urb 1/1
[ 0.011246] [002b73cb] libusb: debug [handle_bulk_completion] last URB in transfer --> complete!
[ 0.011263] [002b73cb] libusb: debug [usbi_handle_transfer_completion] transfer 0x18e6f40 has callback 0x4ad830
[ 0.011308] [002b73cb] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.011318] [002b73c9] libusb: debug [libusb_free_transfer] transfer 0x18e6f40
INFO[0000] read 2:                                     
[ 0.011336] [002b73cb] libusb: debug [handle_events] poll() 3 fds with timeout in 100ms
[ 0.111576] [002b73cb] libusb: debug [handle_events] poll() returned 0
[ 0.111628] [002b73cb] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.111644] [002b73cb] libusb: debug [handle_events] poll() 3 fds with timeout in 100ms
[ 0.111730] [002b73cd] libusb: debug [libusb_alloc_transfer] transfer 0x7f0650000bd0
[ 0.111794] [002b73cd] libusb: debug [libusb_submit_transfer] transfer 0x7f0650000bd0
[ 0.111809] [002b73cd] libusb: debug [submit_bulk_transfer] need 1 urbs for new transfer with length 192
[ 0.111927] [002b73cb] libusb: debug [handle_events] poll() returned 1
[ 0.111951] [002b73cb] libusb: debug [reap_for_handle] urb type=3 status=0 transferred=2
[ 0.111965] [002b73cb] libusb: debug [handle_bulk_completion] handling completion status 0 of bulk urb 1/1
[ 0.111978] [002b73cb] libusb: debug [handle_bulk_completion] last URB in transfer --> complete!
[ 0.111991] [002b73cb] libusb: debug [usbi_handle_transfer_completion] transfer 0x7f0650000bd0 has callback 0x4ad830
[ 0.112034] [002b73cb] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.112047] [002b73cb] libusb: debug [handle_events] poll() 3 fds with timeout in 100ms
[ 0.112058] [002b73cd] libusb: debug [libusb_free_transfer] transfer 0x7f0650000bd0
INFO[0000] read 2:                                     
```

