// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class SourceS3ConfigurationStreamFormatCsvFormatHeaderDefinition
    {
        public readonly Outputs.SourceS3ConfigurationStreamFormatCsvFormatHeaderDefinitionAutogenerated? Autogenerated;
        public readonly Outputs.SourceS3ConfigurationStreamFormatCsvFormatHeaderDefinitionFromCsv? FromCsv;
        public readonly Outputs.SourceS3ConfigurationStreamFormatCsvFormatHeaderDefinitionUserProvided? UserProvided;

        [OutputConstructor]
        private SourceS3ConfigurationStreamFormatCsvFormatHeaderDefinition(
            Outputs.SourceS3ConfigurationStreamFormatCsvFormatHeaderDefinitionAutogenerated? autogenerated,

            Outputs.SourceS3ConfigurationStreamFormatCsvFormatHeaderDefinitionFromCsv? fromCsv,

            Outputs.SourceS3ConfigurationStreamFormatCsvFormatHeaderDefinitionUserProvided? userProvided)
        {
            Autogenerated = autogenerated;
            FromCsv = fromCsv;
            UserProvided = userProvided;
        }
    }
}