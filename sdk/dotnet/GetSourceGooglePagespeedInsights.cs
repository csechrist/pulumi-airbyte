// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceGooglePagespeedInsights
    {
        /// <summary>
        /// SourceGooglePagespeedInsights DataSource
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Airbyte = Pulumi.Airbyte;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySourceGooglepagespeedinsights = Airbyte.GetSourceGooglePagespeedInsights.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSourceGooglePagespeedInsightsResult> InvokeAsync(GetSourceGooglePagespeedInsightsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceGooglePagespeedInsightsResult>("airbyte:index/getSourceGooglePagespeedInsights:getSourceGooglePagespeedInsights", args ?? new GetSourceGooglePagespeedInsightsArgs(), options.WithDefaults());

        /// <summary>
        /// SourceGooglePagespeedInsights DataSource
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Airbyte = Pulumi.Airbyte;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySourceGooglepagespeedinsights = Airbyte.GetSourceGooglePagespeedInsights.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSourceGooglePagespeedInsightsResult> Invoke(GetSourceGooglePagespeedInsightsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceGooglePagespeedInsightsResult>("airbyte:index/getSourceGooglePagespeedInsights:getSourceGooglePagespeedInsights", args ?? new GetSourceGooglePagespeedInsightsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceGooglePagespeedInsightsArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceGooglePagespeedInsightsArgs()
        {
        }
        public static new GetSourceGooglePagespeedInsightsArgs Empty => new GetSourceGooglePagespeedInsightsArgs();
    }

    public sealed class GetSourceGooglePagespeedInsightsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceGooglePagespeedInsightsInvokeArgs()
        {
        }
        public static new GetSourceGooglePagespeedInsightsInvokeArgs Empty => new GetSourceGooglePagespeedInsightsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceGooglePagespeedInsightsResult
    {
        /// <summary>
        /// The values required to configure the source. Parsed as JSON.
        /// </summary>
        public readonly string Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string SourceId;
        public readonly string SourceType;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceGooglePagespeedInsightsResult(
            string configuration,

            string id,

            string name,

            string sourceId,

            string sourceType,

            string workspaceId)
        {
            Configuration = configuration;
            Id = id;
            Name = name;
            SourceId = sourceId;
            SourceType = sourceType;
            WorkspaceId = workspaceId;
        }
    }
}