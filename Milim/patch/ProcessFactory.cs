using HarmonyLib;
using SafeExamBrowser.WindowsApi.Contracts;
using SafeExamBrowser.WindowsApi.Processes;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Milim.patch
{
    [HarmonyPatch(typeof(ProcessFactory), nameof(ProcessFactory.GetAllRunning))]
    public class FilterProcess
    {
        static void Postfix(ref IList<IProcess> __result)
        {

            __result = __result
                .Where(proc => !proc.Name.ToLower().Replace(" ", "").Contains("nava"))
                .Where(proc => !proc.Name.ToLower().Replace(" ", "").Contains("chrome"))
                .Where(proc => !proc.Name.ToLower().Replace(" ", "").Contains("snipping"))
                .Where(proc => !proc.Name.ToLower().Replace(" ", "").Contains("dllhost"))
                .ToList();
        }
    }
}
