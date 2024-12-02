﻿using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Threading.Tasks;
using Dalamud.Game.Gui.PartyFinder.Types;
using Dalamud.Plugin.Services;
using Newtonsoft.Json;

namespace RemotePartyFinder;

internal class Gatherer : IDisposable {
#if DEBUG
    private List<string> UploadURLs = new List<string> { "http://127.0.0.1:8000/contribute/multiple" };
#elif RELEASE
    private List<string> UploadURLs = new List<string> {
        "https://xivpf.com/contribute/multiple",
        "https://findingway.io/receive"
    };
#endif

    private Plugin Plugin { get; }

    private ConcurrentDictionary<int, List<IPartyFinderListing>> Batches { get; } = new();
    private Stopwatch UploadTimer { get; } = new();
    private HttpClient Client { get; } = new();

    internal Gatherer(Plugin plugin) {
        this.Plugin = plugin;

        this.UploadTimer.Start();

        this.Plugin.PartyFinderGui.ReceiveListing += this.OnListing;
        this.Plugin.Framework.Update += this.OnUpdate;
    }

    public void Dispose() {
        this.Plugin.Framework.Update -= this.OnUpdate;
        this.Plugin.PartyFinderGui.ReceiveListing -= this.OnListing;
    }

    private void OnListing(IPartyFinderListing listing, IPartyFinderListingEventArgs args) {
        if (!this.Batches.ContainsKey(args.BatchNumber)) {
            this.Batches[args.BatchNumber] = [];
        }

        this.Batches[args.BatchNumber].Add(listing);
    }

    private void OnUpdate(IFramework framework1) {
        if (this.UploadTimer.Elapsed < TimeSpan.FromSeconds(10)) {
            return;
        }

        this.UploadTimer.Restart();

        foreach (var (batch, listings) in this.Batches.ToList()) {
            this.Batches.Remove(batch, out _);
            Task.Run(async () => {
                var uploadable = listings
                    .Select(listing => new UploadableListing(listing))
                    .ToList();
                var json = JsonConvert.SerializeObject(uploadable);
                foreach (var UploadUrl in UploadURLs)
                {
                    var resp = await this.Client.PostAsync(UploadUrl, new StringContent(json)
                    {
                        Headers = { ContentType = MediaTypeHeaderValue.Parse("application/json") },
                    });
                    var output = await resp.Content.ReadAsStringAsync();
                    Plugin.Log.Info(output);
                }
            });
        }
    }
}
