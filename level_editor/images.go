package main

type LevelImage struct {
	ID   string
	X, Y int
}

var LevelImages = []LevelImage{
	{"small tree", 9032, -794},
	{"huge tree", 8749, -1131},
	{"cave back", 9041, -1065},
	{"small tree", 1420, 424},
	{"big tree", 1173, 309},
	{"square rock", 1226, 395},
	{"big tree", 4251, 184},
	{"huge tree", 4654, -35},
	{"square rock", 4372, 272},
	{"ground left", 7225, -536},
	{"ground center 2", 7280, -536},
	{"ground right", 7398, -535},
	{"ground center 2", 7340, -536},
	{"big tree", 7479, 428},
	{"small tree", 7245, 545},
	{"ground left", 7156, 647},
	{"ground right", 7761, 648},
	{"ground long 2", 7358, 646},
	{"ground long 1", 7209, 646},
	{"big tree", 8369, -905},
	{"small tree", 7800, -789},
	{"huge tree", 7773, -1131},
	{"ground right", 9294, -687},
	{"ground long 1", 8873, -689},
	{"ground long 1", 8672, -689},
	{"ground long 2", 8245, -689},
	{"ground long 2", 7818, -689},
	{"ground left", 7762, -688},
	{"ground long 2", 1071, 532},
	{"ground right", 6911, 319},
	{"ground center 2", 6857, 318},
	{"ground center 2", 6813, 318},
	{"ground center 2", 6753, 318},
	{"ground center 1", 6692, 317},
	{"ground center 3", 6632, 317},
	{"ground left", 6575, 317},
	{"big tree", 6101, 139},
	{"ground right", 6255, 359},
	{"ground center 2", 6195, 357},
	{"ground center 1", 6142, 356},
	{"ground center 1", 6081, 356},
	{"ground left", 6025, 357},
	{"ground right", 5755, 388},
	{"ground center 2", 5702, 387},
	{"ground center 2", 5648, 387},
	{"ground center 1", 5587, 386},
	{"ground left", 5531, 387},
	{"ground center 2", 5243, 409},
	{"ground right", 5289, 410},
	{"ground center 2", 5192, 409},
	{"ground center 1", 5131, 409},
	{"ground center 2", 5010, 409},
	{"ground center 3", 5071, 409},
	{"ground center 1", 4950, 409},
	{"ground center 3", 4891, 409},
	{"ground center 2", 4769, 409},
	{"ground center 2", 4830, 409},
	{"ground center 1", 4647, 409},
	{"ground center 2", 4708, 409},
	{"ground center 3", 4588, 409},
	{"ground center 1", 4526, 409},
	{"ground center 1", 4465, 409},
	{"ground center 1", 4404, 409},
	{"small tree", 5665, 282},
	{"small tree", 5149, 305},
	{"small tree", 4646, 307},
	{"small tree", 3803, 28},
	{"small tree", 1997, 428},
	{"small tree", 3033, 422},
	{"small tree", 2771, 425},
	{"small tree", 640, 427},
	{"small tree", 991, 426},
	{"small tree", 317, 426},
	{"big tree", 5054, 189},
	{"big tree", 3421, -91},
	{"big tree", 2880, 312},
	{"big tree", 1692, 320},
	{"big tree", 775, 318},
	{"huge tree", -6, 89},
	{"huge tree", 1482, 89},
	{"huge tree", 2155, 87},
	{"ground center 2", 1860, 532},
	{"ground center 1", 1677, 532},
	{"ground center 2", 1799, 532},
	{"ground center 1", 3006, 248},
	{"huge tree", -474, 535},
	{"grass left", -355, -75},
	{"grass right", -355, 9},
	{"grass center 1", -358, 82},
	{"grass center 2", -356, 154},
	{"grass center 3", -357, 229},
	{"small tree", -531, 381},
	{"big tree", -330, 290},
	{"ground left", -392, -276},
	{"ground center 1", -243, -285},
	{"ground center 2", -233, -173},
	{"ground center 3", -249, -400},
	{"ground right", -107, -288},
	{"grass center 3", 1682, 529},
	{"grass center 3", 4532, 405},
	{"grass center 2", 5680, 382},
	{"ground right", 3078, 535},
	{"ground center 2", 1921, 532},
	{"ground center 2", 3018, 534},
	{"ground center 1", 2470, 534},
	{"ground center 1", 2531, 534},
	{"ground center 1", 2592, 534},
	{"ground center 1", 2714, 534},
	{"ground center 1", 2896, 534},
	{"ground center 2", 1982, 532},
	{"ground center 2", 2165, 533},
	{"ground center 2", 2287, 534},
	{"ground center 2", 2409, 534},
	{"ground center 2", 2775, 534},
	{"ground center 2", 2835, 534},
	{"ground center 3", 1498, 532},
	{"ground center 3", 1556, 532},
	{"ground center 3", 1739, 532},
	{"ground center 3", 2105, 532},
	{"ground center 3", 2658, 364},
	{"ground center 3", 2654, 534},
	{"ground center 3", 2958, 534},
	{"ground center 1", 1616, 532},
	{"ground center 1", 2043, 532},
	{"ground center 1", 2226, 533},
	{"ground center 1", 2348, 534},
	{"grass center 1", 4426, 405},
	{"grass right", 3086, 527},
	{"grass right", 5296, 403},
	{"grass left", 6015, 353},
	{"ground left", 2576, 364},
	{"ground right", 2714, 365},
	{"ground center 1", 2622, 364},
	{"grass left", 2565, 360},
	{"grass center 1", 2617, 360},
	{"grass right", 2720, 360},
	{"grass center 3", 2668, 360},
	{"ground right", 3045, 250},
	{"ground left", 2911, 249},
	{"ground center 2", 2945, 249},
	{"grass left", 2901, 246},
	{"grass right", 3059, 246},
	{"grass center 3", 2951, 246},
	{"grass center 2", 3004, 246},
	{"grass center 1", 1629, 529},
	{"grass center 1", 1576, 529},
	{"grass center 3", 1523, 529},
	{"grass center 3", 1470, 529},
	{"grass center 3", 3035, 527},
	{"grass center 3", 1947, 529},
	{"grass center 3", 1894, 529},
	{"grass center 3", 1841, 529},
	{"grass center 3", 1735, 529},
	{"grass center 2", 2636, 528},
	{"grass center 2", 1788, 529},
	{"grass center 1", 2689, 528},
	{"grass center 1", 2000, 529},
	{"grass center 3", 2159, 529},
	{"grass center 3", 2106, 529},
	{"grass center 3", 2053, 529},
	{"grass center 3", 2938, 527},
	{"grass center 3", 2740, 528},
	{"grass center 3", 2583, 529},
	{"grass center 3", 2477, 529},
	{"grass center 3", 2265, 529},
	{"grass center 3", 2212, 529},
	{"grass center 3", 2318, 529},
	{"grass center 3", 2371, 529},
	{"grass center 1", 2424, 529},
	{"grass center 3", 2530, 529},
	{"grass center 3", 2791, 528},
	{"grass center 1", 2842, 528},
	{"grass center 3", 2890, 527},
	{"grass center 3", 2986, 527},
	{"ground left", 3389, 132},
	{"ground left", 4226, 409},
	{"ground left", 7329, 290},
	{"ground left", 7829, 439},
	{"ground left", 7654, 49},
	{"ground left", 7265, -160},
	{"ground left", 7709, -354},
	{"ground center 2", 3445, 132},
	{"ground center 1", 3506, 132},
	{"ground center 1", 3567, 132},
	{"ground center 3", 3629, 132},
	{"ground center 2", 3689, 132},
	{"ground center 1", 3750, 132},
	{"ground center 1", 3811, 132},
	{"ground center 1", 3872, 132},
	{"ground center 3", 3934, 132},
	{"ground right", 3993, 133},
	{"grass left", 3382, 127},
	{"grass center 3", 3434, 128},
	{"grass center 3", 3487, 128},
	{"grass center 3", 3540, 128},
	{"grass center 1", 3593, 128},
	{"grass center 2", 3646, 128},
	{"grass center 3", 3699, 129},
	{"grass center 3", 3752, 129},
	{"grass center 2", 3805, 129},
	{"grass center 3", 3857, 129},
	{"grass center 3", 3909, 129},
	{"grass right", 4002, 129},
	{"grass center 2", 3956, 129},
	{"ground center 2", 4282, 409},
	{"ground center 2", 4343, 409},
	{"grass left", 4216, 405},
	{"grass center 3", 4268, 405},
	{"grass center 3", 4321, 405},
	{"grass center 3", 4373, 405},
	{"grass center 1", 5627, 382},
	{"grass center 2", 4635, 404},
	{"grass center 3", 4479, 405},
	{"grass center 3", 6120, 353},
	{"grass center 3", 4584, 404},
	{"grass center 3", 4687, 404},
	{"grass center 3", 4739, 404},
	{"grass center 3", 4792, 404},
	{"grass center 2", 4951, 404},
	{"grass center 1", 4845, 404},
	{"grass center 3", 4898, 404},
	{"grass center 3", 5004, 404},
	{"grass center 3", 5057, 404},
	{"grass center 3", 5110, 404},
	{"grass center 3", 5163, 404},
	{"grass center 2", 5248, 403},
	{"grass center 2", 5212, 403},
	{"grass left", 5522, 382},
	{"grass center 3", 5574, 382},
	{"grass center 1", 6067, 353},
	{"grass center 2", 6221, 353},
	{"grass center 3", 5727, 382},
	{"grass right", 5760, 381},
	{"grass left", 6566, 313},
	{"grass center 3", 6173, 353},
	{"grass center 3", 6617, 313},
	{"grass center 3", 6669, 313},
	{"grass right", 6269, 354},
	{"grass center 3", 6716, 313},
	{"grass center 2", 6767, 313},
	{"grass center 1", 6819, 313},
	{"grass center 1", 6870, 313},
	{"grass right", 6918, 313},
	{"ground long 2", 644, 532},
	{"ground long 1", 217, 532},
	{"ground center 1", 156, 532},
	{"grass long 1", 1152, 529},
	{"grass long 3", 834, 529},
	{"grass long 2", 515, 529},
	{"grass long 2", 197, 529},
	{"grass center 3", 145, 528},
	{"grass left", 7752, -692},
	{"grass long 1", 7804, -692},
	{"grass long 3", 8122, -692},
	{"grass long 2", 8439, -692},
	{"grass long 2", 8757, -692},
	{"grass center 3", 9076, -692},
	{"grass center 3", 9129, -692},
	{"grass center 3", 9182, -692},
	{"grass center 3", 9235, -692},
	{"grass right", 9305, -692},
	{"grass center 2", 9258, -693},
	{"grass long 3", 7200, 643},
	{"grass left", 7150, 643},
	{"grass right", 7776, 643},
	{"grass long 3", 7463, 643},
	{"ground center 1", 7385, 289},
	{"ground center 1", 7445, 289},
	{"ground center 3", 7507, 289},
	{"ground center 1", 7565, 289},
	{"ground right", 7625, 291},
	{"ground right", 7892, 51},
	{"ground center 3", 7711, 49},
	{"ground center 1", 7771, 49},
	{"ground center 1", 7832, 49},
	{"grass left", 7647, 43},
	{"grass center 2", 7350, -542},
	{"grass center 3", 7699, 43},
	{"grass center 3", 7752, 43},
	{"grass center 3", 7805, 44},
	{"grass center 3", 7856, 44},
	{"grass right", 7907, 44},
	{"ground center 3", 7321, -160},
	{"ground center 1", 7399, -160},
	{"ground right", 7459, -159},
	{"ground center 2", 7371, -160},
	{"grass center 3", 7410, -168},
	{"grass right", 7465, -168},
	{"grass center 3", 7357, -168},
	{"grass center 3", 7307, -168},
	{"grass left", 7256, -169},
	{"grass left", 7320, 285},
	{"grass right", 7636, 284},
	{"grass center 3", 7581, 284},
	{"grass center 3", 7529, 284},
	{"grass center 2", 7476, 284},
	{"grass center 3", 7424, 284},
	{"grass center 3", 7371, 285},
	{"ground center 3", 7947, 440},
	{"ground center 1", 7885, 439},
	{"ground center 1", 8006, 440},
	{"ground right", 8047, 441},
	{"grass left", 7821, 432},
	{"grass right", 8052, 433},
	{"grass center 3", 7998, 433},
	{"grass center 3", 7947, 433},
	{"grass center 3", 7896, 433},
	{"grass center 3", 7863, 433},
	{"ground center 2", 7765, -354},
	{"ground center 2", 7826, -354},
	{"ground center 2", 7869, -354},
	{"ground right", 7912, -353},
	{"grass left", 7218, -542},
	{"grass right", 7916, -361},
	{"grass right", 7403, -542},
	{"grass center 3", 7861, -361},
	{"grass center 3", 7808, -361},
	{"grass center 3", 7755, -360},
	{"grass left", 7703, -360},
	{"grass center 3", 7299, -542},
	{"grass center 1", 7264, -542},
	{"cave front", 9041, -1066},
}
