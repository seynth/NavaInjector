using HarmonyLib;
using SafeExamBrowser.WindowsApi.Desktops;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Milim.patch
{
    [HarmonyPatch(typeof(DesktopFactory), "GenerateRandomDesktopName")]
    class GenerateDefaultDesktop
    {
        static void Postfix(ref string __result)
        {
            // tidak akan membuat Desktop baru, memastikan seb jalan di Default desktop
            Console.WriteLine("[Milim::Patch] Stay in Default desktop");
            __result = "Default";
        }


    }

}
