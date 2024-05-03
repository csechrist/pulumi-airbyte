// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceRkiCovid
    {
        /// <summary>
        /// SourceRkiCovid DataSource
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
        ///     var mySourceRkicovid = Airbyte.GetSourceRkiCovid.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSourceRkiCovidResult> InvokeAsync(GetSourceRkiCovidArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceRkiCovidResult>("airbyte:index/getSourceRkiCovid:getSourceRkiCovid", args ?? new GetSourceRkiCovidArgs(), options.WithDefaults());

        /// <summary>
        /// SourceRkiCovid DataSource
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
        ///     var mySourceRkicovid = Airbyte.GetSourceRkiCovid.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSourceRkiCovidResult> Invoke(GetSourceRkiCovidInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceRkiCovidResult>("airbyte:index/getSourceRkiCovid:getSourceRkiCovid", args ?? new GetSourceRkiCovidInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceRkiCovidArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceRkiCovidArgs()
        {
        }
        public static new GetSourceRkiCovidArgs Empty => new GetSourceRkiCovidArgs();
    }

    public sealed class GetSourceRkiCovidInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceRkiCovidInvokeArgs()
        {
        }
        public static new GetSourceRkiCovidInvokeArgs Empty => new GetSourceRkiCovidInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceRkiCovidResult
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
        private GetSourceRkiCovidResult(
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