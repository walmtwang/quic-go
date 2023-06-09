package mocks

//go:generate sh -c "../mockgen_internal.sh mocks tls_extension_handler.go github.com/walmtwang/quic-go/internal/handshake TLSExtensionHandler"
//go:generate sh -c "../mockgen_internal.sh mocks sealer.go github.com/walmtwang/quic-go/internal/handshake Sealer"
//go:generate sh -c "../mockgen_internal.sh mocks stream_flow_controller.go github.com/walmtwang/quic-go/internal/flowcontrol StreamFlowController"
//go:generate sh -c "../mockgen_internal.sh mockackhandler ackhandler/sent_packet_handler.go github.com/walmtwang/quic-go/internal/ackhandler SentPacketHandler"
//go:generate sh -c "../mockgen_internal.sh mockackhandler ackhandler/received_packet_handler.go github.com/walmtwang/quic-go/internal/ackhandler ReceivedPacketHandler"
//go:generate sh -c "../mockgen_internal.sh mocks congestion.go github.com/walmtwang/quic-go/internal/congestion SendAlgorithm"
//go:generate sh -c "../mockgen_internal.sh mocks connection_flow_controller.go github.com/walmtwang/quic-go/internal/flowcontrol ConnectionFlowController"
//go:generate sh -c "../mockgen_internal.sh mockcrypto crypto/aead.go github.com/walmtwang/quic-go/internal/crypto AEAD"
