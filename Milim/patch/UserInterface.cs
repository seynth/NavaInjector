using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using HarmonyLib;
using SafeExamBrowser.UserInterface.Shared.Utilities;

namespace Milim.patch
{
    //bypass window hiding when screencapture
    [HarmonyPatch(typeof(WindowExtensions), nameof(WindowExtensions.ExcludeFromCapture))]
    public class ExcludeFormCapture
    {
        static bool Prefix()
        {
            // tidak menjalankan function ExcludeFromCapture saat screenshot
            return false;
        }
    }
}
