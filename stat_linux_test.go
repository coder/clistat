//go:build linux

package clistat

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestLinuxNumCPU(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name          string
// 		cpuInfo       string
// 		expectedCount int
// 	}{
// 		{
// 			name:          "Raspberry Pi 5 (4 CPU)",
// 			cpuInfo:       procCPUInfoRaspberryPi5,
// 			expectedCount: 4,
// 		},
// 		{
// 			name:          "EU Coder Workspace (96 CPU)",
// 			cpuInfo:       procCPUInfoCoderWorkspaceEU,
// 			expectedCount: 96,
// 		},
// 		{
// 			name:          "UpCloud Cloud Server (1 CPU)",
// 			cpuInfo:       procCPUInfoUpCloud,
// 			expectedCount: 1,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			fs := initFS(t, map[string]string{
// 				"/proc/cpuinfo": tt.cpuInfo,
// 			})

// 			cpuCount := numCPU(fs)
// 			require.Equal(t, tt.expectedCount, cpuCount)
// 		})
// 	}
// }

// const procCPUInfoRaspberryPi5 = `processor       : 0
// BogoMIPS        : 108.00
// Features        : fp asimd evtstrm aes pmull sha1 sha2 crc32 atomics fphp asimdhp cpuid asimdrdm lrcpc dcpop asimddp
// CPU implementer : 0x41
// CPU architecture: 8
// CPU variant     : 0x4
// CPU part        : 0xd0b
// CPU revision    : 1

// processor       : 1
// BogoMIPS        : 108.00
// Features        : fp asimd evtstrm aes pmull sha1 sha2 crc32 atomics fphp asimdhp cpuid asimdrdm lrcpc dcpop asimddp
// CPU implementer : 0x41
// CPU architecture: 8
// CPU variant     : 0x4
// CPU part        : 0xd0b
// CPU revision    : 1

// processor       : 2
// BogoMIPS        : 108.00
// Features        : fp asimd evtstrm aes pmull sha1 sha2 crc32 atomics fphp asimdhp cpuid asimdrdm lrcpc dcpop asimddp
// CPU implementer : 0x41
// CPU architecture: 8
// CPU variant     : 0x4
// CPU part        : 0xd0b
// CPU revision    : 1

// processor       : 3
// BogoMIPS        : 108.00
// Features        : fp asimd evtstrm aes pmull sha1 sha2 crc32 atomics fphp asimdhp cpuid asimdrdm lrcpc dcpop asimddp
// CPU implementer : 0x41
// CPU architecture: 8
// CPU variant     : 0x4
// CPU part        : 0xd0b
// CPU revision    : 1

// Revision        : d04170
// Serial          : 445ad3d19caba6bd
// Model           : Raspberry Pi 5 Model B Rev 1.0`

// const procCPUInfoCoderWorkspaceEU = `processor       : 0
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 0
// cpu cores       : 48
// apicid          : 0
// initial apicid  : 0
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 1
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 1
// cpu cores       : 48
// apicid          : 2
// initial apicid  : 2
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 2
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 2
// cpu cores       : 48
// apicid          : 4
// initial apicid  : 4
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 3
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 3
// cpu cores       : 48
// apicid          : 6
// initial apicid  : 6
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 4
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 4
// cpu cores       : 48
// apicid          : 8
// initial apicid  : 8
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 5
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 5
// cpu cores       : 48
// apicid          : 10
// initial apicid  : 10
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 6
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1519.842
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 32
// cpu cores       : 48
// apicid          : 64
// initial apicid  : 64
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 7
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 33
// cpu cores       : 48
// apicid          : 66
// initial apicid  : 66
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 8
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 34
// cpu cores       : 48
// apicid          : 68
// initial apicid  : 68
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 9
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1519.948
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 35
// cpu cores       : 48
// apicid          : 70
// initial apicid  : 70
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 10
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 36
// cpu cores       : 48
// apicid          : 72
// initial apicid  : 72
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 11
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1518.233
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 37
// cpu cores       : 48
// apicid          : 74
// initial apicid  : 74
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 12
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.962
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 16
// cpu cores       : 48
// apicid          : 32
// initial apicid  : 32
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 13
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1499.735
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 17
// cpu cores       : 48
// apicid          : 34
// initial apicid  : 34
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 14
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1498.968
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 18
// cpu cores       : 48
// apicid          : 36
// initial apicid  : 36
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 15
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 19
// cpu cores       : 48
// apicid          : 38
// initial apicid  : 38
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 16
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 20
// cpu cores       : 48
// apicid          : 40
// initial apicid  : 40
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 17
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.177
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 21
// cpu cores       : 48
// apicid          : 42
// initial apicid  : 42
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 18
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3793.103
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 48
// cpu cores       : 48
// apicid          : 96
// initial apicid  : 96
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 19
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 49
// cpu cores       : 48
// apicid          : 98
// initial apicid  : 98
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 20
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 50
// cpu cores       : 48
// apicid          : 100
// initial apicid  : 100
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 21
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.640
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 51
// cpu cores       : 48
// apicid          : 102
// initial apicid  : 102
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 22
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1499.496
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 52
// cpu cores       : 48
// apicid          : 104
// initial apicid  : 104
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 23
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 53
// cpu cores       : 48
// apicid          : 106
// initial apicid  : 106
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 24
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3798.769
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 24
// cpu cores       : 48
// apicid          : 48
// initial apicid  : 48
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 25
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3255.163
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 25
// cpu cores       : 48
// apicid          : 50
// initial apicid  : 50
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 26
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 26
// cpu cores       : 48
// apicid          : 52
// initial apicid  : 52
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 27
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 27
// cpu cores       : 48
// apicid          : 54
// initial apicid  : 54
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 28
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 28
// cpu cores       : 48
// apicid          : 56
// initial apicid  : 56
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 29
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 29
// cpu cores       : 48
// apicid          : 58
// initial apicid  : 58
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 30
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 2750.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 56
// cpu cores       : 48
// apicid          : 112
// initial apicid  : 112
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 31
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 57
// cpu cores       : 48
// apicid          : 114
// initial apicid  : 114
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 32
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 58
// cpu cores       : 48
// apicid          : 116
// initial apicid  : 116
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 33
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 59
// cpu cores       : 48
// apicid          : 118
// initial apicid  : 118
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 34
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 60
// cpu cores       : 48
// apicid          : 120
// initial apicid  : 120
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 35
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 61
// cpu cores       : 48
// apicid          : 122
// initial apicid  : 122
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 36
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 2750.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 8
// cpu cores       : 48
// apicid          : 16
// initial apicid  : 16
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 37
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 9
// cpu cores       : 48
// apicid          : 18
// initial apicid  : 18
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 38
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 10
// cpu cores       : 48
// apicid          : 20
// initial apicid  : 20
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 39
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 11
// cpu cores       : 48
// apicid          : 22
// initial apicid  : 22
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 40
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 12
// cpu cores       : 48
// apicid          : 24
// initial apicid  : 24
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 41
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3786.332
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 13
// cpu cores       : 48
// apicid          : 26
// initial apicid  : 26
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 42
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3802.923
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 40
// cpu cores       : 48
// apicid          : 80
// initial apicid  : 80
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 43
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 41
// cpu cores       : 48
// apicid          : 82
// initial apicid  : 82
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 44
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 42
// cpu cores       : 48
// apicid          : 84
// initial apicid  : 84
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 45
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 43
// cpu cores       : 48
// apicid          : 86
// initial apicid  : 86
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 46
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 2750.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 44
// cpu cores       : 48
// apicid          : 88
// initial apicid  : 88
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 47
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 45
// cpu cores       : 48
// apicid          : 90
// initial apicid  : 90
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 48
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 0
// cpu cores       : 48
// apicid          : 1
// initial apicid  : 1
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 49
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1519.799
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 1
// cpu cores       : 48
// apicid          : 3
// initial apicid  : 3
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 50
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 2
// cpu cores       : 48
// apicid          : 5
// initial apicid  : 5
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 51
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 3
// cpu cores       : 48
// apicid          : 7
// initial apicid  : 7
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 52
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3793.990
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 4
// cpu cores       : 48
// apicid          : 9
// initial apicid  : 9
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 53
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 5
// cpu cores       : 48
// apicid          : 11
// initial apicid  : 11
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 54
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1518.852
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 32
// cpu cores       : 48
// apicid          : 65
// initial apicid  : 65
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 55
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 33
// cpu cores       : 48
// apicid          : 67
// initial apicid  : 67
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 56
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 34
// cpu cores       : 48
// apicid          : 69
// initial apicid  : 69
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 57
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 35
// cpu cores       : 48
// apicid          : 71
// initial apicid  : 71
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 58
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3799.946
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 36
// cpu cores       : 48
// apicid          : 73
// initial apicid  : 73
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 59
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 37
// cpu cores       : 48
// apicid          : 75
// initial apicid  : 75
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 60
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1498.004
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 16
// cpu cores       : 48
// apicid          : 33
// initial apicid  : 33
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 61
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 17
// cpu cores       : 48
// apicid          : 35
// initial apicid  : 35
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 62
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 18
// cpu cores       : 48
// apicid          : 37
// initial apicid  : 37
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 63
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 19
// cpu cores       : 48
// apicid          : 39
// initial apicid  : 39
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 64
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 20
// cpu cores       : 48
// apicid          : 41
// initial apicid  : 41
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 65
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 21
// cpu cores       : 48
// apicid          : 43
// initial apicid  : 43
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 66
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 48
// cpu cores       : 48
// apicid          : 97
// initial apicid  : 97
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 67
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 49
// cpu cores       : 48
// apicid          : 99
// initial apicid  : 99
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 68
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 50
// cpu cores       : 48
// apicid          : 101
// initial apicid  : 101
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 69
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 51
// cpu cores       : 48
// apicid          : 103
// initial apicid  : 103
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 70
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 52
// cpu cores       : 48
// apicid          : 105
// initial apicid  : 105
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 71
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 53
// cpu cores       : 48
// apicid          : 107
// initial apicid  : 107
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 72
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 2750.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 24
// cpu cores       : 48
// apicid          : 49
// initial apicid  : 49
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 73
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3794.665
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 25
// cpu cores       : 48
// apicid          : 51
// initial apicid  : 51
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 74
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3799.923
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 26
// cpu cores       : 48
// apicid          : 53
// initial apicid  : 53
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 75
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 27
// cpu cores       : 48
// apicid          : 55
// initial apicid  : 55
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 76
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 28
// cpu cores       : 48
// apicid          : 57
// initial apicid  : 57
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 77
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 29
// cpu cores       : 48
// apicid          : 59
// initial apicid  : 59
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 78
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 56
// cpu cores       : 48
// apicid          : 113
// initial apicid  : 113
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 79
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3801.845
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 57
// cpu cores       : 48
// apicid          : 115
// initial apicid  : 115
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 80
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1514.193
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 58
// cpu cores       : 48
// apicid          : 117
// initial apicid  : 117
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 81
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 59
// cpu cores       : 48
// apicid          : 119
// initial apicid  : 119
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 82
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 60
// cpu cores       : 48
// apicid          : 121
// initial apicid  : 121
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 83
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 61
// cpu cores       : 48
// apicid          : 123
// initial apicid  : 123
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 84
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 3798.093
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 8
// cpu cores       : 48
// apicid          : 17
// initial apicid  : 17
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 85
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 2750.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 9
// cpu cores       : 48
// apicid          : 19
// initial apicid  : 19
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 86
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 10
// cpu cores       : 48
// apicid          : 21
// initial apicid  : 21
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 87
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 11
// cpu cores       : 48
// apicid          : 23
// initial apicid  : 23
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 88
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 12
// cpu cores       : 48
// apicid          : 25
// initial apicid  : 25
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 89
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 13
// cpu cores       : 48
// apicid          : 27
// initial apicid  : 27
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 90
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 40
// cpu cores       : 48
// apicid          : 81
// initial apicid  : 81
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 91
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 41
// cpu cores       : 48
// apicid          : 83
// initial apicid  : 83
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 92
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1519.908
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 42
// cpu cores       : 48
// apicid          : 85
// initial apicid  : 85
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 93
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 43
// cpu cores       : 48
// apicid          : 87
// initial apicid  : 87
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 94
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 44
// cpu cores       : 48
// apicid          : 89
// initial apicid  : 89
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]

// processor       : 95
// vendor_id       : AuthenticAMD
// cpu family      : 25
// model           : 17
// model name      : AMD EPYC 9454P 48-Core Processor
// stepping        : 1
// microcode       : 0xa101148
// cpu MHz         : 1500.000
// cache size      : 1024 KB
// physical id     : 0
// siblings        : 96
// core id         : 45
// cpu cores       : 48
// apicid          : 91
// initial apicid  : 91
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc rep_good amd_lbr_v2 nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba perfmon_v2 ibrs ibpb stibp ibrs_enhanced vmmcall fsgsbase bmi1 avx2 smep bmi2 erms invpcid cqm rdt_a avx512f avx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local user_shstk avx512_bf16 clzero irperf xsaveerptr rdpru wbnoinvd amd_ppin cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshold avic v_vmsave_vmload vgif x2avic v_spec_ctrl vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq la57 rdpid overflow_recov succor smca fsrm flush_l1d debug_swap
// bugs            : sysret_ss_attrs spectre_v1 spectre_v2 spec_store_bypass srso
// bogomips        : 5499.87
// TLB size        : 3584 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 52 bits physical, 57 bits virtual
// power management: ts ttp tm hwpstate cpb eff_freq_ro [13] [14]`

// const procCPUInfoUpCloud = `processor       : 0
// vendor_id       : AuthenticAMD
// cpu family      : 23
// model           : 49
// model name      : AMD EPYC 7542 32-Core Processor
// stepping        : 0
// microcode       : 0x1000065
// cpu MHz         : 2894.560
// cache size      : 512 KB
// physical id     : 0
// siblings        : 1
// core id         : 0
// cpu cores       : 1
// apicid          : 0
// initial apicid  : 0
// fpu             : yes
// fpu_exception   : yes
// cpuid level     : 16
// wp              : yes
// flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm rep_good nopl cpuid extd_apicid tsc_known_freq pni pclmulqdq ssse3 fma cx16 sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm cmp_legacy cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw perfctr_core ssbd ibrs ibpb stibp vmmcall fsgsbase tsc_adjust bmi1 avx2 smep bmi2 rdseed adx smap clflushopt clwb sha_ni xsaveopt xsavec xgetbv1 clzero xsaveerptr wbnoinvd arat umip rdpid arch_capabilities
// bugs            : sysret_ss_attrs null_seg spectre_v1 spectre_v2 spec_store_bypass retbleed smt_rsb srso
// bogomips        : 5789.12
// TLB size        : 1024 4K pages
// clflush size    : 64
// cache_alignment : 64
// address sizes   : 48 bits physical, 48 bits virtual
// power management:`
